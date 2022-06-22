package fcm

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"firebase.google.com/go/v4/messaging"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/option"
)

const (
	iidEndpoint = "https://iid.googleapis.com/iid/v1"
)

func testTopic(t *testing.T, isSubscribe bool) {
	httpmock.ActivateNonDefault(httpClient)
	defer httpmock.DeactivateAndReset()

	endpoint := iidEndpoint
	if isSubscribe {
		endpoint += ":batchAdd"
	} else {
		endpoint += ":batchRemove"
	}

	var b []byte
	httpmock.RegisterResponder(http.MethodPost, endpoint, func(r *http.Request) (resp *http.Response, err error) {
		b, _ = io.ReadAll(r.Body)
		var any map[string]interface{}
		byteExpect := []byte("{\"results\": [{}, {\"error\": \"error_reason\"}]}")
		_ = json.Unmarshal(byteExpect, &any)

		resp, err = httpmock.NewJsonResponse(http.StatusOK, any)
		return
	})

	ctx := context.Background()
	client, err := newClient(ctx, nil, option.WithHTTPClient(httpClient))
	fatalIfError(t, err)

	var resp *messaging.TopicManagementResponse
	if isSubscribe {
		resp, err = client.SubscribeToTopic(ctx, []string{"id1", "id2"}, "test-topic")
	} else {
		resp, err = client.UnsubscribeFromTopic(ctx, []string{"id1", "id2"}, "test-topic")
	}
	fatalIfError(t, err)

	var parsed map[string]interface{}
	err = json.Unmarshal(b, &parsed)
	fatalIfError(t, err)

	want := map[string]interface{}{
		"to":                  "/topics/test-topic",
		"registration_tokens": []interface{}{"id1", "id2"},
	}
	assert.EqualValues(t, want, parsed)
	assert.Equal(t, 1, resp.SuccessCount)
	assert.Equal(t, 1, resp.FailureCount)
	assert.Equal(t, 1, len(resp.Errors))
	assert.Equal(t, 1, resp.Errors[0].Index)
	assert.Equal(t, "error_reason", resp.Errors[0].Reason)
}

func TestSubscribe(t *testing.T) {
	testTopic(t, true)
}

func TestUnsubscribe(t *testing.T) {
	testTopic(t, false)
}

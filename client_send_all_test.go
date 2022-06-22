package fcm

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"firebase.google.com/go/v4/messaging"
	"github.com/stretchr/testify/assert"
)

var testSuccessResponse = []fcmResponse{
	{
		Name: "projects/test-project/messages/1",
	},
	{
		Name: "projects/test-project/messages/2",
	},
}
var testMessages = []*messaging.Message{
	{Topic: "topic1"},
	{Topic: "topic2"},
}

const wantMime = "multipart/mixed; boundary=__END_OF_PART__"

func testSendAll(t *testing.T, dryRun bool) {
	resp, err := createMultipartResponse(testSuccessResponse, nil)
	fatalIfError(t, err)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", wantMime)
		_, _ = w.Write(resp)
	}))
	defer ts.Close()

	ctx := context.Background()
	client, err := newClient(ctx, ts)
	fatalIfError(t, err)

	var br *messaging.BatchResponse
	if dryRun {
		br, err = client.SendAllDryRun(ctx, testMessages)
	} else {
		br, err = client.SendAll(ctx, testMessages)
	}
	fatalIfError(t, err)

	assert.Equal(t, 2, br.SuccessCount)
	assert.Equal(t, 0, br.FailureCount)
	assert.Equal(t, 2, len(br.Responses))

	for idx, r := range br.Responses {
		assert.True(t, r.Success)
		assert.Nil(t, r.Error)
		assert.Equal(t, testSuccessResponse[idx].Name, r.MessageID)
	}

}

func TestSendAll(t *testing.T) {
	testSendAll(t, false)
}

func TestSendAllDryRun(t *testing.T) {
	testSendAll(t, true)
}

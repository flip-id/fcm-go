package fcm

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"firebase.google.com/go/v4/messaging"
	"github.com/stretchr/testify/assert"
)

var testMulticastMessage = &messaging.MulticastMessage{
	Tokens: []string{"token1", "token2"},
}

func testSendMulticast(t *testing.T, dryRun bool) {
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
		br, err = client.SendMulticastDryRun(ctx, testMulticastMessage)
	} else {
		br, err = client.SendMulticast(ctx, testMulticastMessage)
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

func TestSendMulticast(t *testing.T) {
	testSendMulticast(t, false)
}

func TestSendMulticastDryRun(t *testing.T) {
	testSendMulticast(t, true)
}

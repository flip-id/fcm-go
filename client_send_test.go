package fcm

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"firebase.google.com/go/v4/messaging"
	"github.com/stretchr/testify/assert"
)

func testSend(t *testing.T, dryRun bool) {
	var b []byte
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("{ \"name\":\"" + testMessageID + "\" }"))
	}))
	defer ts.Close()

	ctx := context.Background()
	c, err := newClient(ctx, ts)
	fatalIfError(t, err)

	var (
		validMessages = []struct {
			name string
			req  *messaging.Message
			want map[string]interface{}
		}{
			{
				name: "NotificationMessage",
				req: &messaging.Message{
					Notification: &messaging.Notification{
						Title:    "t",
						Body:     "b",
						ImageURL: "http://image.jpg",
					},
					Topic: "test-topic",
				},
				want: map[string]interface{}{
					"notification": map[string]interface{}{
						"title": "t",
						"body":  "b",
						"image": "http://image.jpg",
					},
					"topic": "test-topic",
				},
			},
		}
		res string
	)
	if dryRun {
		res, err = c.SendDryRun(ctx, validMessages[0].req)
	} else {
		res, err = c.Send(ctx, validMessages[0].req)
	}
	fatalIfError(t, err)

	assert.Equal(t, testMessageID, res)

	var parsed map[string]interface{}
	err = json.Unmarshal(b, &parsed)
	fatalIfError(t, err)

	assert.EqualValues(t, validMessages[0].want, parsed["message"])

}

func TestSend(t *testing.T) {
	testSend(t, false)
}

func TestSendDryRun(t *testing.T) {
	testSend(t, true)
}

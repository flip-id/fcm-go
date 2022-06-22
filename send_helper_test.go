package fcm

import (
	"context"
	"net/http/httptest"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
)

type MockTokenSource struct {
	AccessToken string
}

func (ts *MockTokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{AccessToken: ts.AccessToken}, nil
}

const testMessageID = "projects/test-project/messages/msg_id"

func defaultFirebaseConfig() *firebase.Config {
	return &firebase.Config{
		ProjectID: "test-project",
	}
}

func defaultFirebaseOption() []option.ClientOption {
	return []option.ClientOption{
		option.WithTokenSource(&MockTokenSource{AccessToken: "test-token"}),
	}
}

func newMessagingClient(ctx context.Context, ts *httptest.Server, opts ...option.ClientOption) (c *messaging.Client, err error) {
	var url string
	if ts != nil {
		url = ts.URL
	}

	app, err := firebase.NewApp(
		ctx,
		defaultFirebaseConfig(),
		append(defaultFirebaseOption(), append(opts, option.WithEndpoint(url))...)...,
	)
	if err != nil {
		return
	}

	c, err = app.Messaging(ctx)
	return
}

func newClient(ctx context.Context, ts *httptest.Server, opts ...option.ClientOption) (c *client, err error) {
	msgClient, err := newMessagingClient(ctx, ts, opts...)
	if err != nil {
		return
	}

	c = (&client{Client: msgClient}).Assign(new(Option).Init())
	return
}

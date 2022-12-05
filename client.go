package fcm

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/fairyhunter13/reflecthelper/v4"
)

//go:generate mockgen -package=fcm -source=client.go -destination=client_mock.go

// Client is an interface to interact with the FCM client.
type Client interface {
	// Send : inherit doc from https://pkg.go.dev/firebase.google.com/go/messaging#Client.Send.
	Send(ctx context.Context, message *messaging.Message) (string, error)
	// SendAll : inherit doc from https://pkg.go.dev/firebase.google.com/go/messaging#Client.SendAll.
	SendAll(ctx context.Context, messages []*messaging.Message) (*messaging.BatchResponse, error)
	// SendAllDryRun: inherit doc from https://pkg.go.dev/firebase.google.com/go/messaging#Client.SendAllDryRun.
	SendAllDryRun(ctx context.Context, messages []*messaging.Message) (*messaging.BatchResponse, error)
	// SendDryRun : inherit doc from https://pkg.go.dev/firebase.google.com/go/messaging#Client.SendDryRun.
	SendDryRun(ctx context.Context, message *messaging.Message) (string, error)
	// SendMulticast : inherit doc from https://pkg.go.dev/firebase.google.com/go/messaging#Client.SendMulticast.
	SendMulticast(ctx context.Context, message *messaging.MulticastMessage) (*messaging.BatchResponse, error)
	// SendMulticastDryRun : inherit doc from https://pkg.go.dev/firebase.google.com/go/messaging#Client.SendMulticastDryRun.
	SendMulticastDryRun(ctx context.Context, message *messaging.MulticastMessage) (*messaging.BatchResponse, error)
	// SubscribeToTopic : inherit doc from https://pkg.go.dev/firebase.google.com/go/messaging#Client.SubscribeToTopic.
	SubscribeToTopic(ctx context.Context, tokens []string, topic string) (*messaging.TopicManagementResponse, error)
	// UnsubscribeFromTopic : inherit doc from https://pkg.go.dev/firebase.google.com/go/messaging#Client.UnsubscribeFromTopic.
	UnsubscribeFromTopic(ctx context.Context, tokens []string, topic string) (*messaging.TopicManagementResponse, error)
}

type client struct {
	*messaging.Client
	opt *Option
}

func (c *client) Assign(o *Option) *client {
	c.opt = reflecthelper.CloneInterface(o).(*Option)
	return c
}

// New returns a new client wrapping the Firebase messaging client.
func New(ctx context.Context, opts ...FnOption) (c Client, err error) {
	o := new(Option)
	for _, opt := range opts {
		opt(o)
	}
	o.Init()

	app, err := firebase.NewApp(ctx, o.Config, o.Options...)
	if err != nil {
		return
	}

	msg, err := app.Messaging(ctx)
	if err != nil {
		return
	}

	c = (&client{
		Client: msg,
	}).Assign(o)
	return
}

// Send : inherit doc from https://pkg.go.dev/firebase.google.com/go/messaging#Client.Send.
func (c *client) Send(ctx context.Context, message *messaging.Message) (res string, err error) {
	s := newSession().
		SetCtx(ctx).
		SetMessage(message)
	c.opt.HookBefore(s)
	defer func() {
		c.opt.HookAfter(s.SetMessageID(res).SetError(err))
	}()

	res, err = c.Client.Send(s.Ctx, s.Message)
	return
}

// SendAll : inherit doc from https://pkg.go.dev/firebase.google.com/go/messaging#Client.SendAll.
func (c *client) SendAll(ctx context.Context, messages []*messaging.Message) (msg *messaging.BatchResponse, err error) {
	s := newSession().
		SetCtx(ctx).
		SetMessages(messages)
	c.opt.HookBefore(s)
	defer func() {
		c.opt.HookAfter(s.SetBatchResponse(msg).SetError(err))
	}()

	msg, err = c.Client.SendAll(s.Ctx, s.Messages)
	return
}

// SendAllDryRun: inherit doc from https://pkg.go.dev/firebase.google.com/go/messaging#Client.SendAllDryRun.
func (c *client) SendAllDryRun(ctx context.Context, messages []*messaging.Message) (msg *messaging.BatchResponse, err error) {
	s := newSession().
		SetCtx(ctx).
		SetMessages(messages)
	c.opt.HookBefore(s)
	defer func() {
		c.opt.HookAfter(s.SetBatchResponse(msg).SetError(err))
	}()

	msg, err = c.Client.SendAllDryRun(s.Ctx, s.Messages)
	return
}

// SendDryRun : inherit doc from https://pkg.go.dev/firebase.google.com/go/messaging#Client.SendDryRun.
func (c *client) SendDryRun(ctx context.Context, message *messaging.Message) (res string, err error) {
	s := newSession().
		SetCtx(ctx).
		SetMessage(message)
	c.opt.HookBefore(s)
	defer func() {
		c.opt.HookAfter(s.SetMessageID(res).SetError(err))
	}()

	res, err = c.Client.SendDryRun(s.Ctx, s.Message)
	return
}

// SendMulticast : inherit doc from https://pkg.go.dev/firebase.google.com/go/messaging#Client.SendMulticast.
func (c *client) SendMulticast(ctx context.Context, message *messaging.MulticastMessage) (msg *messaging.BatchResponse, err error) {
	s := newSession().
		SetCtx(ctx).
		SetMulticastMessage(message)
	c.opt.HookBefore(s)
	defer func() {
		c.opt.HookAfter(s.SetBatchResponse(msg).SetError(err))
	}()

	msg, err = c.Client.SendMulticast(s.Ctx, s.MulticastMessage)
	return
}

// SendMulticastDryRun : inherit doc from https://pkg.go.dev/firebase.google.com/go/messaging#Client.SendMulticastDryRun.
func (c *client) SendMulticastDryRun(ctx context.Context, message *messaging.MulticastMessage) (msg *messaging.BatchResponse, err error) {
	s := newSession().
		SetCtx(ctx).
		SetMulticastMessage(message)
	c.opt.HookBefore(s)
	defer func() {
		c.opt.HookAfter(s.SetBatchResponse(msg).SetError(err))
	}()

	msg, err = c.Client.SendMulticastDryRun(s.Ctx, s.MulticastMessage)
	return
}

// SubscribeToTopic : inherit doc from https://pkg.go.dev/firebase.google.com/go/messaging#Client.SubscribeToTopic.
func (c *client) SubscribeToTopic(ctx context.Context, tokens []string, topic string) (msg *messaging.TopicManagementResponse, err error) {
	s := newSession().
		SetCtx(ctx).
		SetTokens(tokens).
		SetTopic(topic)
	c.opt.HookBefore(s)
	defer func() {
		c.opt.HookAfter(s.SetTopicManagementResponse(msg).SetError(err))
	}()

	msg, err = c.Client.SubscribeToTopic(s.Ctx, s.Tokens, s.Topic)
	return
}

// UnsubscribeFromTopic : inherit doc from https://pkg.go.dev/firebase.google.com/go/messaging#Client.UnsubscribeFromTopic.
func (c *client) UnsubscribeFromTopic(ctx context.Context, tokens []string, topic string) (msg *messaging.TopicManagementResponse, err error) {
	s := newSession().
		SetCtx(ctx).
		SetTokens(tokens).
		SetTopic(topic)
	c.opt.HookBefore(s)
	defer func() {
		c.opt.HookAfter(s.SetTopicManagementResponse(msg).SetError(err))
	}()

	msg, err = c.Client.UnsubscribeFromTopic(s.Ctx, s.Tokens, s.Topic)
	return
}

package fcm

import (
	"context"

	"firebase.google.com/go/v4/messaging"
)

// Session is the collection of variables for each request.
type Session struct {
	// Arguments
	Ctx              context.Context
	Message          *messaging.Message
	Messages         []*messaging.Message
	MulticastMessage *messaging.MulticastMessage
	Tokens           []string
	Topic            string

	// Returns
	MessageID               string
	Error                   error
	BatchResponse           *messaging.BatchResponse
	TopicManagementResponse *messaging.TopicManagementResponse
}

func (s *Session) SetCtx(ctx context.Context) *Session {
	s.Ctx = ctx
	return s
}

func (s *Session) SetMessage(msg *messaging.Message) *Session {
	s.Message = msg
	return s
}

func (s *Session) SetMessages(msgs []*messaging.Message) *Session {
	s.Messages = msgs
	return s
}

func (s *Session) SetMulticastMessage(msg *messaging.MulticastMessage) *Session {
	s.MulticastMessage = msg
	return s
}

func (s *Session) SetTokens(tokens []string) *Session {
	s.Tokens = tokens
	return s
}

func (s *Session) SetTopic(topic string) *Session {
	s.Topic = topic
	return s
}

func (s *Session) SetMessageID(msgID string) *Session {
	s.MessageID = msgID
	return s
}

func (s *Session) SetError(err error) *Session {
	s.Error = err
	return s
}

func (s *Session) SetBatchResponse(resp *messaging.BatchResponse) *Session {
	s.BatchResponse = resp
	return s
}

func (s *Session) SetTopicManagementResponse(resp *messaging.TopicManagementResponse) *Session {
	s.TopicManagementResponse = resp
	return s
}

var noopHook = func(s *Session) {}

// FnHook is a function for executing hook in actions.
type FnHook func(s *Session)

func newSession() *Session {
	return new(Session)
}

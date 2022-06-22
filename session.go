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

// SetCtx sets the context of the Session.
func (s *Session) SetCtx(ctx context.Context) *Session {
	s.Ctx = ctx
	return s
}

// SetMessage sets the single message of the Session.
func (s *Session) SetMessage(msg *messaging.Message) *Session {
	s.Message = msg
	return s
}

// SetMessages sets the batch messages of the Session.
func (s *Session) SetMessages(msgs []*messaging.Message) *Session {
	s.Messages = msgs
	return s
}

// SetMulticastMessage sets the multicast message of the Session.
func (s *Session) SetMulticastMessage(msg *messaging.MulticastMessage) *Session {
	s.MulticastMessage = msg
	return s
}

// SetTokens sets the batch tokens of the Session.
func (s *Session) SetTokens(tokens []string) *Session {
	s.Tokens = tokens
	return s
}

// SetTopic sets the topic of the Session.
func (s *Session) SetTopic(topic string) *Session {
	s.Topic = topic
	return s
}

// SetMessageID sets the message ID of the Session from the return arguments.
func (s *Session) SetMessageID(msgID string) *Session {
	s.MessageID = msgID
	return s
}

// SetError sets the error of the Session from the return arguments.
func (s *Session) SetError(err error) *Session {
	s.Error = err
	return s
}

// SetBatchResponse sets the btach response of the Session from the return arguments.
func (s *Session) SetBatchResponse(resp *messaging.BatchResponse) *Session {
	s.BatchResponse = resp
	return s
}

// SetTopicManagementResponse sets the topic management response of the Session from the return arguments.
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

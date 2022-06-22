package fcm

import (
	"context"
	"log"
	"time"

	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type contextKey string

func ExampleNew() {
	ctx := context.Background()
	fnBefore := func(s *Session) {
		dummyData := map[string][]string{
			"test": {"test-header-1"},
		}
		var span ddtrace.Span
		span, s.Ctx = tracer.StartSpanFromContext(
			s.Ctx,
			"http.out",
			tracer.SpanType(ext.SpanTypeHTTP),
			tracer.ServiceName("fcm-go"),
			tracer.Measured(),
			tracer.ResourceName("send"),
		)

		s.Ctx = context.WithValue(s.Ctx, contextKey("begin_time"), time.Now())
		if err := tracer.Inject(span.Context(), tracer.HTTPHeadersCarrier(dummyData)); err != nil {
			log.Printf("Error injecting span context to fcm-go client! Error: %v", err)
		}
	}
	fnAfter := func(s *Session) {
		span, _ := tracer.SpanFromContext(s.Ctx)
		defer func() {
			span.Finish(tracer.WithError(s.Error))
		}()

		begin := s.Ctx.Value(contextKey("begin_time")).(time.Time)
		log.Println(
			"span_id", span.Context().SpanID(),
			"trace_id", span.Context().TraceID(),
			"type", "send",
			"duration", time.Since(begin),
			"message_id", s.MessageID,
		)
		if s.Error != nil {
			log.Printf("Fcm-go client gets error: %v", s.Error)
		}
	}
	fcmClient, err := New(ctx,
		WithAfterHook(fnAfter),
		WithBeforeHook(fnBefore),
		WithClientOptions(
			option.WithCredentialsFile("test/data/service_account.json"),
		),
	)
	if err != nil {
		log.Fatalf("Error in getting the new client: %v", err)
	}

	msgID, err := fcmClient.Send(ctx, &messaging.Message{})
	if err != nil {
		log.Fatalf("Error in sending the message: %v", err)
	}

	log.Printf("The message ID: %s", msgID)
}

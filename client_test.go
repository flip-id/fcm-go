package fcm

import (
	"context"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/api/option"
)

func TestNew(t *testing.T) {
	type args struct {
		ctx    context.Context
		opts   []FnOption
		server *httptest.Server
	}
	tests := []struct {
		name          string
		args          func() args
		wantNilClient bool
		wantErr       bool
	}{
		{
			name: "valid client",
			args: func() args {

				ctx := context.Background()
				return args{
					opts: []FnOption{
						WithConfig(nil),
						WithAfterHook(noopHook),
						WithBeforeHook(noopHook),
						WithClientOptions(
							option.WithCredentialsFile("test/data/service_account.json"),
						),
					},
					ctx: ctx,
				}
			},
			wantNilClient: false,
			wantErr:       false,
		},
		{
			name: "invalid firebase client",
			args: func() args {
				s, cfg, err := getNewAppParams()
				fatalIfError(t, err)

				os.Setenv(FirebaseEnvName, "testrandom.json")
				ctx := context.Background()
				return args{
					server: s,
					opts: []FnOption{
						WithClientOptions(
							option.WithTokenSource(cfg.TokenSource(ctx)),
						),
					},
					ctx: ctx,
				}
			},
			wantNilClient: true,
			wantErr:       true,
		},
		{
			name: "invalid messaging client",
			args: func() args {
				s, cfg, err := getNewAppParams()
				fatalIfError(t, err)

				os.Setenv(FirebaseEnvName, "")
				ctx := context.Background()
				return args{
					server: s,
					opts: []FnOption{
						WithClientOptions(
							option.WithTokenSource(cfg.TokenSource(ctx)),
						),
					},
					ctx: ctx,
				}
			},
			wantNilClient: true,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.args()
			if args.server != nil {
				defer args.server.Close()
			}

			gotC, err := New(args.ctx, args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantNilClient {
				assert.Nil(t, gotC, "Client should be nil!")
			}
		})
	}
}

package fcm

import (
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

// Option is a colletion of options for the client.
type Option struct {
	Config     *firebase.Config
	Options    []option.ClientOption
	HookBefore FnHook
	HookAfter  FnHook
}

// Init initializes the Option struct to avoid nil panic.
func (o *Option) Init() *Option {
	if o.HookBefore == nil {
		o.HookBefore = noopHook
	}

	if o.HookAfter == nil {
		o.HookAfter = noopHook
	}

	return o
}

// FnOption is a functional option to set the Option.
type FnOption func(o *Option)

// WithConfig sets the Config of the Option.
func WithConfig(cfg *firebase.Config) FnOption {
	return func(o *Option) {
		o.Config = cfg
	}
}

// WithClientOptions sets the client options to the Option.
func WithClientOptions(opts ...option.ClientOption) FnOption {
	return func(o *Option) {
		o.Options = opts
	}
}

// WithBeforeHook sets the hook for before action.
func WithBeforeHook(fn FnHook) FnOption {
	return func(o *Option) {
		o.HookBefore = fn
	}
}

// WithAfterHook sets the hook for after action.
func WithAfterHook(fn FnHook) FnOption {
	return func(o *Option) {
		o.HookAfter = fn
	}
}

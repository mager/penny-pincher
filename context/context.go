package context

import "context"

func ProvideContext() context.Context {
	return context.Background()
}

var Options = ProvideContext

package openrpc

import "context"

type MethodDesc struct {
	Name        string
	Description string
	Handler     methodHandler
}

type methodHandler func(ctx *context.Context, server interface{}) (interface{}, error)

type ServiceDesc struct {
	Name string
	// Check whether the implementation satisfies the interface.
	HandlerType interface{}
	Description string
}

type ParameterDesc struct {
	Name        string
	Description string
}

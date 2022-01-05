package openrpc

import (
	"context"
	"reflect"
)

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

type Server struct {
	// The service descriptor.
	ServiceDesc ServiceDesc
	// The method descriptors.
	MethodDescs []MethodDesc
}

func (s *Server) Serve(ctx context.Context, server interface{}) error {
	return nil
}

func (s *Server) RegisterService(sd *ServiceDesc, impl interface{}) error {
	if impl != nil {
		s.ServiceDesc = *sd
	}
	ht := reflect.TypeOf(sd.HandlerType).Elem()
	st := reflect.TypeOf(impl)
    if !st.Implements(ht) {
	return nil
}

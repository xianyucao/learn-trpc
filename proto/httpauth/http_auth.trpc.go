// Code generated by trpc-go/trpc-cmdline v1.0.5. DO NOT EDIT.
// source: http_auth.proto

package httpauth

import (
	"context"
	"errors"
	"fmt"

	_ "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/codec"
	_ "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/server"
)

// START ======================================= Server Service Definition ======================================= START

// AuthService defines service.
type AuthService interface {
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) // @alias=/demo/auth/Login
}

func AuthService_Login_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &LoginRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(AuthService).Login(ctx, reqbody.(*LoginRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// AuthServer_ServiceDesc descriptor for server.RegisterService.
var AuthServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "demo.httpauth.Auth",
	HandlerType: ((*AuthService)(nil)),
	Methods: []server.Method{
		{
			Name: "/demo/auth/Login",
			Func: AuthService_Login_Handler,
		},
		{
			Name: "/demo.httpauth.Auth/Login",
			Func: AuthService_Login_Handler,
		},
	},
}

// RegisterAuthService registers service.
func RegisterAuthService(s server.Service, svr AuthService) {
	if err := s.Register(&AuthServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("Auth register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedAuth struct{}

func (s *UnimplementedAuth) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, errors.New("rpc Login of service Auth is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// AuthClientProxy defines service client proxy
type AuthClientProxy interface {
	Login(ctx context.Context, req *LoginRequest, opts ...client.Option) (rsp *LoginResponse, err error) // @alias=/demo/auth/Login
}

type AuthClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewAuthClientProxy = func(opts ...client.Option) AuthClientProxy {
	return &AuthClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *AuthClientProxyImpl) Login(ctx context.Context, req *LoginRequest, opts ...client.Option) (*LoginResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/demo/auth/Login")
	msg.WithCalleeServiceName(AuthServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("Auth")
	msg.WithCalleeMethod("Login")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &LoginResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END

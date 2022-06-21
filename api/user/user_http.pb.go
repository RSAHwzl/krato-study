// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package user

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserServiceCreateUser = "/api.user.UserService/CreateUser"
const OperationUserServiceGetUserList = "/api.user.UserService/GetUserList"
const OperationUserServiceSetUserCache = "/api.user.UserService/SetUserCache"
const OperationUserServiceGetUserCache = "/api.user.UserService/GetUserCache"

type UserServiceHTTPServer interface {
	CreateUser(context.Context, *CreateUserReq) (*CreateUserRes, error)
	GetUserCache(context.Context, *GetUserCacheReq) (*GetUserCacheRes, error)
	GetUserList(context.Context, *GetUserListReq) (*GetUserListRes, error)
	SetUserCache(context.Context, *SetUserCacheReq) (*SetUserCacheRes, error)
}

func RegisterUserServiceHTTPServer(s *http.Server, srv UserServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/createUser", _UserService_CreateUser0_HTTP_Handler(srv))
	r.GET("/api/getUserList", _UserService_GetUserList0_HTTP_Handler(srv))
	r.POST("/api/setUserCache", _UserService_SetUserCache0_HTTP_Handler(srv))
	r.GET("/api/getUserCache/{name}", _UserService_GetUserCache0_HTTP_Handler(srv))
}

func _UserService_CreateUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserRes)
		return ctx.Result(200, reply)
	}
}

func _UserService_GetUserList0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserListReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceGetUserList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserList(ctx, req.(*GetUserListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserListRes)
		return ctx.Result(200, reply)
	}
}

func _UserService_SetUserCache0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetUserCacheReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceSetUserCache)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetUserCache(ctx, req.(*SetUserCacheReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetUserCacheRes)
		return ctx.Result(200, reply)
	}
}

func _UserService_GetUserCache0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserCacheReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceGetUserCache)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserCache(ctx, req.(*GetUserCacheReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserCacheRes)
		return ctx.Result(200, reply)
	}
}

type UserServiceHTTPClient interface {
	CreateUser(ctx context.Context, req *CreateUserReq, opts ...http.CallOption) (rsp *CreateUserRes, err error)
	GetUserCache(ctx context.Context, req *GetUserCacheReq, opts ...http.CallOption) (rsp *GetUserCacheRes, err error)
	GetUserList(ctx context.Context, req *GetUserListReq, opts ...http.CallOption) (rsp *GetUserListRes, err error)
	SetUserCache(ctx context.Context, req *SetUserCacheReq, opts ...http.CallOption) (rsp *SetUserCacheRes, err error)
}

type UserServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewUserServiceHTTPClient(client *http.Client) UserServiceHTTPClient {
	return &UserServiceHTTPClientImpl{client}
}

func (c *UserServiceHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserReq, opts ...http.CallOption) (*CreateUserRes, error) {
	var out CreateUserRes
	pattern := "/api/createUser"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) GetUserCache(ctx context.Context, in *GetUserCacheReq, opts ...http.CallOption) (*GetUserCacheRes, error) {
	var out GetUserCacheRes
	pattern := "/api/getUserCache/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceGetUserCache))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) GetUserList(ctx context.Context, in *GetUserListReq, opts ...http.CallOption) (*GetUserListRes, error) {
	var out GetUserListRes
	pattern := "/api/getUserList"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceGetUserList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserServiceHTTPClientImpl) SetUserCache(ctx context.Context, in *SetUserCacheReq, opts ...http.CallOption) (*SetUserCacheRes, error) {
	var out SetUserCacheRes
	pattern := "/api/setUserCache"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceSetUserCache))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

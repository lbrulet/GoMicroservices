// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/users/users.proto

package go_micro_srv_users

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Users service

type UsersService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*UserResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*UserResponse, error)
}

type usersService struct {
	c    client.Client
	name string
}

func NewUsersService(name string, c client.Client) UsersService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.users"
	}
	return &usersService{
		c:    c,
		name: name,
	}
}

func (c *usersService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "Users.Create", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "Users.Login", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersHandler interface {
	Create(context.Context, *CreateRequest, *UserResponse) error
	Login(context.Context, *LoginRequest, *UserResponse) error
}

func RegisterUsersHandler(s server.Server, hdlr UsersHandler, opts ...server.HandlerOption) error {
	type users interface {
		Create(ctx context.Context, in *CreateRequest, out *UserResponse) error
		Login(ctx context.Context, in *LoginRequest, out *UserResponse) error
	}
	type Users struct {
		users
	}
	h := &usersHandler{hdlr}
	return s.Handle(s.NewHandler(&Users{h}, opts...))
}

type usersHandler struct {
	UsersHandler
}

func (h *usersHandler) Create(ctx context.Context, in *CreateRequest, out *UserResponse) error {
	return h.UsersHandler.Create(ctx, in, out)
}

func (h *usersHandler) Login(ctx context.Context, in *LoginRequest, out *UserResponse) error {
	return h.UsersHandler.Login(ctx, in, out)
}
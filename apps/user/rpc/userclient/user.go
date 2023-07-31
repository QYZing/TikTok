// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"TikTok/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BackgroundImageReq  = user.BackgroundImageReq
	BackgroundImageResp = user.BackgroundImageResp
	BasicUserInfo       = user.BasicUserInfo
	BasicUserInfoReq    = user.BasicUserInfoReq
	BasicUserInfoResp   = user.BasicUserInfoResp
	LoginReq            = user.LoginReq
	LoginResp           = user.LoginResp
	RegisterReq         = user.RegisterReq
	RegisterResp        = user.RegisterResp
	SetAvatarReq        = user.SetAvatarReq
	SetAvatarResp       = user.SetAvatarResp

	User interface {
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		Detail(ctx context.Context, in *BasicUserInfoReq, opts ...grpc.CallOption) (*BasicUserInfoResp, error)
		SetAvatar(ctx context.Context, in *SetAvatarReq, opts ...grpc.CallOption) (*SetAvatarResp, error)
		SetBackgroundImage(ctx context.Context, in *BackgroundImageReq, opts ...grpc.CallOption) (*BackgroundImageResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) Detail(ctx context.Context, in *BasicUserInfoReq, opts ...grpc.CallOption) (*BasicUserInfoResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Detail(ctx, in, opts...)
}

func (m *defaultUser) SetAvatar(ctx context.Context, in *SetAvatarReq, opts ...grpc.CallOption) (*SetAvatarResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.SetAvatar(ctx, in, opts...)
}

func (m *defaultUser) SetBackgroundImage(ctx context.Context, in *BackgroundImageReq, opts ...grpc.CallOption) (*BackgroundImageResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.SetBackgroundImage(ctx, in, opts...)
}
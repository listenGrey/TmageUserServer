package grpc

import (
	"github.com/listenGrey/TmagegRpcPKG/userInfo"
	"google.golang.org/grpc/peer"
	"log"

	"context"
)

type RegisterExistenceServer struct {
	userInfo.UnimplementedCheckExistenceServer
}

func (u *RegisterExistenceServer) RegisterCheck(ctx context.Context, email *userInfo.RegisterEmail) (*userInfo.Existence, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("检查该用户是否存在")
	}

	// 检查邮箱
	if email.GetEmail() == "chen.luo@lixil.com" {
		return &userInfo.Existence{Exsit: true}, nil
	} else {
		return &userInfo.Existence{Exsit: false}, nil
	}
}

type RegisterServer struct {
	userInfo.UnimplementedRegisterInfoServer
}

func (r *RegisterServer) Register(ctx context.Context, form *userInfo.RegisterForm) (*userInfo.Success, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("用户注册")
	}

	return &userInfo.Success{Success: true}, nil
}

type LoginServer struct {
	userInfo.UnimplementedLoginCheckServer
}

func (u *LoginServer) LoginCheck(ctx context.Context, user *userInfo.LoginForm) (*userInfo.LogInfo, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("用户登录")
	}

	return &userInfo.LogInfo{
		UserID: 111,
		Info:   1000,
	}, nil
}

package grpc

import (
	"TmageUsersServer/dao"
	"github.com/listenGrey/TmagegRpcPKG/userInfo"
	"google.golang.org/grpc/peer"
	"log"

	"context"
)

type ExistenceServer struct {
	userInfo.UnimplementedCheckExistenceServer
}

func (u *ExistenceServer) RegisterCheck(ctx context.Context, email *userInfo.RegisterEmail) (*userInfo.Existence, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("检查该用户是否存在")
	}

	// 检查邮箱
	return &userInfo.Existence{Exsit: dao.CheckEmail(email.GetEmail())}, nil

}

type RegisterServer struct {
	userInfo.UnimplementedRegisterInfoServer
}

func (r *RegisterServer) Register(ctx context.Context, form *userInfo.RegisterForm) (*userInfo.Success, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("用户注册")
	}

	return &userInfo.Success{Success: dao.InsertUser(form)}, nil
}

type LoginServer struct {
	userInfo.UnimplementedLoginCheckServer
}

func (u *LoginServer) LoginCheck(ctx context.Context, user *userInfo.LoginForm) (*userInfo.LogInfo, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("用户登录")
	}
	info, userID := dao.Login(user)
	return &userInfo.LogInfo{Info: info, UserID: userID}, nil
}

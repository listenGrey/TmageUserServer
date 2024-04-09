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
	exist, info := dao.CheckEmail(email.GetEmail())
	return &userInfo.Existence{Exist: exist, Info: info}, nil

}

type RegisterServer struct {
	userInfo.UnimplementedRegisterInfoServer
}

func (r *RegisterServer) Register(ctx context.Context, form *userInfo.RegisterForm) (*userInfo.Success, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("用户注册")
	}
	flag, info := dao.InsertUser(form)
	return &userInfo.Success{Success: flag, Info: info}, nil
}

type LoginServer struct {
	userInfo.UnimplementedLoginCheckServer
}

func (u *LoginServer) LoginCheck(ctx context.Context, user *userInfo.LoginForm) (*userInfo.LogInfo, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("用户登录")
	}
	userID, info := dao.Login(user)
	return &userInfo.LogInfo{UserID: userID, Info: info}, nil
}

package grpc

import (
	"TmageUsersServer/models"
	"github.com/listenGrey/TmagegRpcPKG/userInfo"
)

func RegisterFormMarshal(r *userInfo.RegisterForm) models.Register {
	var user models.Register

	user.Email = r.GetEmail()
	user.UserID = r.GetUserID()
	user.UserName = r.GetUserName()
	user.Password = r.GetPassword()

	return user
}

func LoginMarshal(form *userInfo.LoginForm) models.Login {
	var login models.Login

	login.Email = form.GetEmail()
	login.Password = form.GetPassword()

	return login
}

package models

import (
	"github.com/listenGrey/TmagegRpcPKG/userInfo"
)

func RegisterFormMarshal(r *userInfo.RegisterForm) Register {
	var user Register

	user.Email = r.GetEmail()
	user.UserID = r.GetUserID()
	user.UserName = r.GetUserName()
	user.Password = r.GetPassword()

	return user
}

func LoginMarshal(form *userInfo.LoginForm) Login {
	var login Login

	login.Email = form.GetEmail()
	login.Password = form.GetPassword()

	return login
}

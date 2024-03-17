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

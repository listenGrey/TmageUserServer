package models

type User struct {
	UserID   uint64 `json:"user_id" bson:"user_id"`
	UserName string `json:"userName" bson:"user_name"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}

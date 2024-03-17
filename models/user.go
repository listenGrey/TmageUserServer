package models

type Register struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
	UserName string `bson:"user_name"`
	UserID   int64  `bson:"user_id"`
}

type Login struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type LogInfo struct {
	Info   int64
	UserID int64
}

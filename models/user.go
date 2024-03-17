package models

type Register struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
	UserName string `bson:"user_name"`
	UserID   int64  `bson:"_id"`
}

type Login struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

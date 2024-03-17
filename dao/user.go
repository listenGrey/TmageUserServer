package dao

import (
	"TmageUsersServer/models"
	"context"
	"github.com/listenGrey/TmagegRpcPKG/userInfo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Code int64

const (
	StatusSuccess      Code = 1000
	StatusUserNotExist Code = 1003
	StatusInvalidPwd   Code = 1004
	StatusBusy         Code = 1005
	StatusConnDBErr    Code = 1202
)

func CheckEmail(email string) bool {
	// 连接DB
	client := MongoDBClient("tmageUser", "user")
	if client == nil {
		log.Printf("无法连接到MongoDB")
		return false
	}

	// 创建一个过滤器
	filter := bson.D{{"email", bson.D{{"$eq", email}}}}

	// 查找文档
	cursor, err := client.CountDocuments(context.Background(), filter)
	if err != nil {
		err.Error()
		return false
	}

	// 返回结果
	if cursor == 0 {
		return false
	} else {
		return true
	}

}

func InsertUser(user *userInfo.RegisterForm) bool {
	// 连接DB
	client := MongoDBClient("tmageUser", "user")
	if client == nil {
		log.Printf("无法连接到MongoDB")
		return false
	}
	// 插入数据
	userData, err := bson.Marshal(models.RegisterFormMarshal(user))
	_, err = client.InsertOne(context.TODO(), userData)
	if err != nil {
		err.Error()
		return false
	}
	return true
}

func Login(user *userInfo.LoginForm) (info, userID int64) {
	// 连接DB
	client := MongoDBClient("tmageUser", "user")
	if client == nil {
		log.Printf("无法连接到MongoDB")
		info = int64(StatusConnDBErr)
		userID = 0
		return
	}

	// 检查email是否存在
	filter := bson.D{{"email", user.GetEmail()}}

	var loginUser models.Register
	err := client.FindOne(context.Background(), filter).Decode(&loginUser)
	if err == mongo.ErrNoDocuments {
		info = int64(StatusUserNotExist)
		userID = 0
		return
	} else if err != nil {
		err.Error()
		info = int64(StatusBusy)
		userID = 0
		return
	}

	// 检查密码是否匹配
	if loginUser.Password != user.GetPassword() {
		info = int64(StatusInvalidPwd)
		userID = 0
		return
	}

	info = int64(StatusSuccess)
	userID = loginUser.UserID
	return

}

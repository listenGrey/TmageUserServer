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
	StatusConnDBERR    Code = 1104
)

func CheckEmail(email string) (bool, int64) {
	// 连接DB
	client := MongoDBClient("tmageUser", "user")
	if client == nil {
		log.Printf("无法连接到MongoDB")
		return false, int64(StatusConnDBERR)
	}

	// 创建一个过滤器
	filter := bson.D{{"email", bson.D{{"$eq", email}}}}

	// 查找文档
	cursor, err := client.CountDocuments(context.Background(), filter)
	if err != nil {
		err.Error()
		return false, int64(StatusBusy)
	}

	// 返回结果
	if cursor == 0 {
		return false, int64(StatusSuccess)
	} else {
		return true, int64(StatusSuccess)
	}

}

func InsertUser(user *userInfo.RegisterForm) (bool, int64) {
	// 连接DB
	client := MongoDBClient("tmageUser", "user")
	if client == nil {
		log.Printf("无法连接到MongoDB")
		return false, int64(StatusConnDBERR)
	}
	// 插入数据
	userData, err := bson.Marshal(models.RegisterFormMarshal(user))
	_, err = client.InsertOne(context.TODO(), userData)
	if err != nil {
		err.Error()
		return false, int64(StatusBusy)
	}
	return true, int64(StatusSuccess)
}

func Login(user *userInfo.LoginForm) (userID, info int64) {
	// 连接DB
	client := MongoDBClient("tmageUser", "user")
	if client == nil {
		log.Printf("无法连接到MongoDB")
		userID = 0
		info = int64(StatusConnDBERR)
		return
	}

	// 检查email是否存在
	filter := bson.D{{"email", user.GetEmail()}}

	var loginUser models.Register
	err := client.FindOne(context.Background(), filter).Decode(&loginUser)
	if err == mongo.ErrNoDocuments {
		userID = 0
		info = int64(StatusUserNotExist)
		return
	} else if err != nil {
		err.Error()
		userID = 0
		info = int64(StatusBusy)
		return
	}

	// 检查密码是否匹配
	if loginUser.Password != user.GetPassword() {
		userID = 0
		info = int64(StatusInvalidPwd)
		return
	}

	userID = loginUser.UserID
	info = int64(StatusSuccess)
	return

}

package dao

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"
)

func MongoDBClient(db, col string) *mongo.Collection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		err.Error()
		return nil
	}

	// 选择数据库和集合
	database := client.Database(db)
	collection := database.Collection(col)

	return collection
}

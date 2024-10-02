package user_module

import (
	"context"
	app_config "eslam3400/go-backend-api/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Collection

func Init() {
	db = app_config.DB.Collection("users")
}

func CreateUser(user SUser) error {
	_, err := db.InsertOne(context.TODO(), user)
	return err
}

func GetUserByID(id string) (SUser, error) {
	var user SUser
	err := db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(user)
	return user, err
}

func GetAllUsers() ([]SUser, error) {
	var users []SUser
	cursor, err := db.Find(context.TODO(), bson.M{})
	if err != nil {
		return users, err
	}
	defer cursor.Close(context.Background())
	cursor.All(context.Background(), &users)
	return users, nil
}

package app_config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBClient *mongo.Client
var DB *mongo.Database

func ConnectToDB() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(Env.DBHost).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	DBClient = client
	DB = client.Database(Env.DBName)
	fmt.Println(`
		Connected to DB Successfully! 
		***************************
	`)
	fmt.Println()
}

func DisconnectDB() {
	err := DBClient.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Disconnected from DB")
}

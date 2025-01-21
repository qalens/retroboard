package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var Db *mongo.Database

func SetupConnection() {
	var err error
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URI"))
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	Db = client.Database("retroboard")
}

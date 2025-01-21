package migrations

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"qalens.com/retroboard/pkg/models"
	"qalens.com/retroboard/pkg/models/db"
)

func AllMigrations() {
	{
		db.Db.CreateCollection(context.Background(), models.UsersCollection)
		indexModel := mongo.IndexModel{
			Keys:    bson.D{{Key: "username", Value: 1}}, // Create index on `username` field
			Options: options.Index().SetUnique(true),     // Ensure the index is unique
		}
		db.Db.Collection(models.UsersCollection).Indexes().CreateOne(context.Background(), indexModel)
	}
}

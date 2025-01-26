package models

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"qalens.com/retroboard/pkg/models/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const UsersCollection = "users"

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
func CreateUser(username string, password string) (*User, error) {
	if hashedPassword, err := hashPassword(password); err == nil {
		user := &User{
			Username:  username,
			Password:  hashedPassword,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		collection := db.Db.Collection(UsersCollection)
		insertResult, err := collection.InsertOne(context.Background(), user) // Using the struct directly
		if err != nil {
			return nil, err
		}
		user.Id = insertResult.InsertedID.(primitive.ObjectID)
		return user, nil
	} else {
		return nil, err
	}
}
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func Authenticate(username string, password string) (*User, error) {
	var dbUser User
	collection := db.Db.Collection(UsersCollection)
	err := collection.FindOne(context.Background(), bson.D{{Key: "username", Value: username}}).Decode(&dbUser) // Using the struct for FindOne
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}
	if !checkPasswordHash(password, dbUser.Password) {
		return nil, fmt.Errorf("invalid credentials")
	} else {
		return &User{
			Id:        dbUser.Id,
			Username:  dbUser.Username,
			CreatedAt: dbUser.CreatedAt,
			UpdatedAt: dbUser.UpdatedAt,
		}, nil
	}

}

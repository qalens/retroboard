package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"qalens.com/retroboard/pkg/models"
)

func generateJWT(user *models.User) (string, error) {
	// Create the Claims
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
		Issuer:    "retroboard",
		Subject:   user.Id.String(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func Login(username string, password string) (string, error) {
	if user, err := models.Authenticate(username, password); err == nil {
		return generateJWT(user)
	} else {
		return "", err
	}
}
func SignUp(username string, password string) (string, error) {
	if user, err := models.CreateUser(username, password); err == nil {
		return generateJWT(user)
	} else {
		return "", err
	}
}
func Authenticate(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, fmt.Errorf("not authenticated")
	}
	claims := token.Claims.(jwt.StandardClaims)
	return &claims, nil
}

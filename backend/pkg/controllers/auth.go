package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"qalens.com/retroboard/pkg/services"
)

type AuthController struct{}
type CredentialsPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ac AuthController) SignUp(ctx *gin.Context) {
	var credentials CredentialsPayload
	ctx.ShouldBindBodyWithJSON(&credentials)
	if token, err := services.SignUp(credentials.Username, credentials.Password); err == nil {
		ctx.SetCookie("token", token, 3600*24, "/", os.Getenv("DOMAIN"), false, true)
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Signup success",
			"data":    token,
		})
		return
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Signup fail",
			"data":    err.Error(),
		})
	}
}
func (ac AuthController) Login(ctx *gin.Context) {
	var credentials CredentialsPayload
	ctx.ShouldBindBodyWithJSON(&credentials)
	if token, err := services.Login(credentials.Username, credentials.Password); err == nil {
		ctx.SetCookie("token", token, 3600*24, "/", os.Getenv("DOMAIN"), false, true)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Login success",
			"data":    token,
		})

		return
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "login failed",
			"data":    err.Error(),
		})
	}
}
func (ac AuthController) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", os.Getenv("DOMAIN"), false, true)
	ctx.Redirect(http.StatusTemporaryRedirect, "/home")
}

package pkg

import (
	"github.com/gin-gonic/gin"
	"qalens.com/retroboard/pkg/controllers"
)

func SetupRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		authCtrl := controllers.AuthController{}
		auth.POST("/signup", authCtrl.SignUp)
		auth.POST("/login", authCtrl.Login)
		auth.GET("/logout", authCtrl.Logout)
	}
}

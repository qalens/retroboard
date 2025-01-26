package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"qalens.com/retroboard/pkg"
	"qalens.com/retroboard/pkg/models/db"
	"qalens.com/retroboard/pkg/models/db/migrations"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.SetupConnection()
	migrations.AllMigrations()
}
func main() {

	router := gin.Default()
	pkg.SetupRoutes(router)
	router.Run(":8080")
}

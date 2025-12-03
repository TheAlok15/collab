package main

import (
	// "fmt"
	"log"

	"github.com/TheAlok15/collab/internal/database"
	"github.com/TheAlok15/collab/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
if err != nil {
    log.Println("No .env file found")
}

	r := gin.Default()

	database.ConnectDB()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello Alok Go backend running!"})
		c.JSON(501, gin.H{"status code" : 100})
	})
	r.POST("/signup",handlers.Signup)
	r.POST("/sigin",handlers.Signin)

	r.Run(":8080")
}
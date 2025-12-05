package main

import (
	// "fmt"
	"log"
	"time"

	"github.com/TheAlok15/collab/internal/database"
	"github.com/TheAlok15/collab/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	r := gin.Default()

	database.ConnectDB()
	r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:5173"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    MaxAge:           12 * time.Hour,
}))

r.OPTIONS("/*path", func(c *gin.Context) {
    c.Status(200)
})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello Alok Go backend running!"})
		// c.JSON(501, gin.H{"status code": 100})
	})
	r.POST("/signup", handlers.Signup)
	r.POST("/signin", handlers.Signin)

	r.Run(":8080")
}

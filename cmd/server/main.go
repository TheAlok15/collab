package main

import(
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/TheAlok15/collab/internal/database"
)

func main(){
	r := gin.Default()

	database.ConnectDB()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello Alok Go backend running!"})
		c.JSON(501, gin.H{"status code" : 100})
	})

	r.Run(":8080")
}
package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pusher/pusher-http-go/v5"
)

func main() {
	pusherClient := pusher.Client{
		AppID:   "1744075",
		Key:     "8810ee0117e1eeeece9d",
		Secret:  "390bb649bed1933138a8",
		Cluster: "mt1",
		Secure:  true,
	}
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Adjust with your React app's URL
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/api/chat", func(c *gin.Context) {
		var data map[string]string

		if err := c.ShouldBindJSON(&data); err != nil {
			fmt.Println(err.Error())
		}
		err := pusherClient.Trigger("chat", "message", data)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.Run(":8080") // Adjust with your desired port
}

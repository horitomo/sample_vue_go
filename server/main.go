package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type clientValue struct {
	Msg string `json:"msg"`
}

type clientLoginValue struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	r.POST("/api", func(c *gin.Context) {
		var client clientValue
		c.BindJSON(&client)
		log.Println(client.Msg)
		log.Println("hogehoge")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"message": "ping",
			"request": client.Msg,
		})
	})

	r.POST("/login", func(c *gin.Context) {
		var loginForm clientLoginValue
		c.BindJSON(&loginForm)
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"message": "ping",
			"request": loginForm,
		})
	})

	r.Run(":8888")
}

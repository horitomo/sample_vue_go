package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

type clientValue struct {
	Msg string `json:"msg"`
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
    r.Run(":8888")
}
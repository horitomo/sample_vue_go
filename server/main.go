package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type clientValue struct {
	Msg string `json:"msg"`
}

type clientLoginValue struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Users ユーザー情報を持つ構造体
type Users struct {
	ID       int
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "go"
	PASS := "golang"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "golang"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

func main() {
	var usersSession Users

	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	defer db.Close()

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
		// c.Header("Access-Control-Allow-Origin", "*")
		// c.JSON(200, gin.H{
		// 	"message": "ping",
		// 	"request": loginForm,
		// })

		var result Users
		error := db.Where("email = ? AND password = ?", loginForm.Email, loginForm.Password).Find(&result).Error
		if error != nil {
			return
		}
		log.Println(result)
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"request": result,
		})
	})

	r.POST("/register", func(c *gin.Context) {
		var msg string
		var users Users
		c.BindJSON(&users)
		log.Println(users)
		error := db.Create(&users).Error
		if error != nil {
			msg = "aaa"
			fmt.Println(error)
		} else {
			fmt.Println("データ追加成功")
		}
		usersSession = users
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"error":   msg,
			"request": usersSession,
		})
	})

	r.Run(":8888")
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //ไม่ตั้งชื่อให้โดยใช้ _
	"github.com/jmoiron/sqlx"
)

//config file
const (
	username = "root"
	password = "hello1234"
	host     = "localhost"
	port     = "3306"
	database = "helloworld"
)

// Hello is contain Name GLOBAL VALUE JAA IF PRIVATE IS name
type Hello struct {
	Name string `json:"name"`
}

// Register model
type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User model
type User struct {
	UID      string `db:"uid"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func main() {
	urlSql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)
	_, err := sqlx.Connect("mysql", urlSql)
	if err != nil {
		log.Println(err)
		log.Fatal("Can not Connet Database")
	}
	//route
	route := gin.Default()
	route.GET("/hello", helloworldHandler)     //END POINT JAA
	route.POST("/helloname", helloNameHandler) //END POINT JAA
	route.POST("/register", registerHandler)   //END POINT JAA
	route.Run(":3000")
}

func helloworldHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func helloNameHandler(context *gin.Context) {
	var helloName Hello
	err := context.ShouldBindJSON(&helloName)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "input incorrect",
		})
		return
	}
	helloName.Name = "hello " + helloName.Name
	context.JSON(http.StatusOK, helloName)
}

func registerHandler(context *gin.Context) {
	var helloName Hello
	err := context.ShouldBindJSON(&helloName)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "input incorrect",
		})
		return
	}
	helloName.Name = "hello " + helloName.Name
	context.JSON(http.StatusOK, helloName)
}

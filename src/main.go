package main

import (
	"fmt"
	"helloworld/api"
	"helloworld/repository"
	"log"

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

func main() {
	urlSQL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)
	db, err := sqlx.Connect("mysql", urlSQL)
	if err != nil {
		log.Println(err)
		log.Fatal("Can not Connet Database")
	}
	userRepo := repository.UserRepoMYSQL{
		DBConnection: db,
	}
	userAPI := api.UserAPI{
		UserRepo: &userRepo,
	}

	//route
	route := gin.Default()
	route.GET("/hello", api.HelloworldHandler)       //END POINT JAA
	route.POST("/helloname", api.HelloNameHandler)   //END POINT JAA
	route.POST("/register", userAPI.RegisterHandler) //END POINT JAA
	route.Run(":3000")
}

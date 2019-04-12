package main

import (
	"fmt"
	"helloworld/model"
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

// UserRepo ที่เกี่ยวกับดาต้าเบส
type UserRepo struct {
	DBConnection *sqlx.DB
}

func main() {
	urlSQL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)
	db, err := sqlx.Connect("mysql", urlSQL)
	if err != nil {
		log.Println(err)
		log.Fatal("Can not Connet Database")
	}
	userRepo := UserRepo{
		DBConnection: db,
	}

	//route
	route := gin.Default()
	route.GET("/hello", helloworldHandler)            //END POINT JAA
	route.POST("/helloname", helloNameHandler)        //END POINT JAA
	route.POST("/register", userRepo.registerHandler) //END POINT JAA
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

//registerHandler มี method ของ UserRepo
func (userRepo UserRepo) registerHandler(context *gin.Context) {
	var user model.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "input incorrect",
		})
		return
	}
	if err := createUser(user, userRepo.DBConnection); err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
	}
	context.Status(http.StatusCreated)
}

func createUser(user model.User, DBConnection *sqlx.DB) error {
	statement := `INSERT INTO user (username, password) VALUES (?,?)`
	tx := DBConnection.MustBegin()
	tx.MustExec(statement, user.Username, user.Password)
	if err := tx.Commit(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

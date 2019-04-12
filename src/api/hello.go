package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//model
// Hello is contain Name GLOBAL VALUE JAA IF PRIVATE IS name
type Hello struct {
	Name string `json:"name"`
}

func HelloworldHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func HelloNameHandler(context *gin.Context) {
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

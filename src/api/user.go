package api

import (
	"helloworld/model"
	"helloworld/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserAPI
type UserAPI struct {
	UserRepo repository.UserRepo
}

//registerHandler มี method ของ UserAPI
func (userAPI UserAPI) RegisterHandler(context *gin.Context) {
	var user model.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "input incorrect",
		})
		return
	}
	if err := userAPI.UserRepo.CreateUser(user); err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
	}
	context.Status(http.StatusCreated)
}

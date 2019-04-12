package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()
	route.GET("/hello", helloworld)
	route.Run(":3000")
}

func helloworld(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "hello world",
	})
}

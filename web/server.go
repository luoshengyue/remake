package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "everything is ok",
		})
	})
	err := r.Run()
	if err != nil {
		fmt.Println("something wrong with gin")
	}
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello Word",
		})
	})
	r.POST("/", func(ctx *gin.Context) {
		var user User
		err := ctx.BindJSON(&user)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, user)
	})

	r.Run(":8080")
}

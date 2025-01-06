package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name" binding:"required"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details"`
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
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse{
				Error:   "Invalid request",
				Details: "The 'name' field is required and must be valid JSON",
			})
			return
		}

		ctx.JSON(http.StatusOK, user)
	})

	r.Run(":8080")
}

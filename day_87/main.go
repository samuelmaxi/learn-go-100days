package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/webhook", func(ctx *gin.Context) {
		var payload map[string]interface{}
		if err := ctx.BindJSON(&payload); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Recebido: ", payload)
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run(":3000")

}

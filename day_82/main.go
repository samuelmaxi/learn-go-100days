package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		log.WithFields(logrus.Fields{
			"status":     ctx.Writer.Status(),
			"method":     ctx.Request.Method,
			"path":       ctx.Request.URL.Path,
			"client_ip":  ctx.ClientIP(),
			"latency":    time.Since(start),
			"user_agent": ctx.Request.UserAgent(),
		}).Info("Requisição recevida")
	})

	r.GET("/", func(ctx *gin.Context) {
		log.Println(http.StatusOK)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "helo word",
		})
	})

	r.Run()
}

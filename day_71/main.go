package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var users = map[string]string{
	"test": "123",
}

var oaut2Config = oauth2.Config{
	ClientID:     "test-client-id",
	ClientSecret: "test-client-secret",
	RedirectURL:  "http://localhost:8080/callback",
	Scopes:       []string{"read", "write"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "http://localhost:8080/auth",
		TokenURL: "http://localhost:8080/token",
	},
}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		username := ctx.DefaultQuery("username", "")
		password := ctx.DefaultQuery("password", "")

		if pass, exists := users[username]; exists && pass == password {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Login Successful",
			})
			ctx.Redirect(http.StatusFound, oaut2Config.RedirectURL+"?code=authorization_code")
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid credentials",
			})
		}
	})

	r.GET("/auth", func(ctx *gin.Context) {
		code := ctx.DefaultQuery("code", "")
		if code == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Authorization code missing",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Authorization successful",
			"code":    code,
		})
	})

	r.GET("/token", func(ctx *gin.Context) {
		code := ctx.DefaultQuery("code", "")
		if code == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Authorization code missing",
			})
			return
		}

		token := oauth2.Token{
			AccessToken:  "access_token",
			TokenType:    "bearer",
			RefreshToken: "refresh_token",
			Expiry:       time.Now().Add(1 * time.Hour),
		}

		ctx.JSON(http.StatusOK, gin.H{
			"access_token":  token.AccessToken,
			"token_type":    token.TokenType,
			"expires_in":    int(token.Expiry.Sub(time.Now()).Seconds()),
			"refresh_token": token.RefreshToken,
		})
	})

	r.Run(":8080")

}

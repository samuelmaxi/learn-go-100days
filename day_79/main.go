package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func startMockServer() {
	r := gin.Default()

	r.GET("/api/users", func(ctx *gin.Context) {
		users := []User{
			{ID: 1, Name: "Alice", Email: "alice@email.com"},
			{ID: 2, Name: "Bob", Email: "bob@email.com"},
		}
		ctx.JSON(http.StatusOK, users)
	})
	r.Run(":8082")
}

func fetchUsers() {
	resp, err := http.Get("http://localhost:8082/api/users")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var users []User
	json.Unmarshal(body, &users)

	fmt.Println("Usuários recebidos da API externa:")
	for _, user := range users {
		fmt.Printf(" - %s (%s)\n", user.Name, user.Email)
	}
}

func main() {
	go startMockServer()
	fmt.Println("Aguardando moc server")

	fetchUsers()
}

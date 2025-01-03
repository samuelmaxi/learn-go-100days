package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users []User

func main() {

	http.HandleFunc("/create", handlerCreate)
	http.HandleFunc("/read", handlerRead)
	http.HandleFunc("/updated", handlerUpdated)
	http.HandleFunc("/deleted", handlerDelete)

	fmt.Println(users)

	slog.Info("Servidor rodando na porta :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// CREATE
func handlerCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		req, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var u User

		err = json.Unmarshal(req, &u)
		if err != nil {
			panic(err)
		}
		fmt.Println("RETORNO DE BODY", u)
		w.Write([]byte(u.Name))

		users = append(users, u)
		fmt.Println(users)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`Metodo nao encontrado`))
}

// READ
func handlerRead(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		jsonData, err := json.Marshal(users)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`Metodo nao encontrado`))
}

// UPDATED
func handlerUpdated(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		query := r.URL.Query()
		param := query.Get("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			panic(err)
		}

		fmt.Println("PASSOU")
		fmt.Println("id: ", id)

		for i := 0; i < len(users); i++ {
			fmt.Println("id do usuario", users[i].ID)
			if users[i].ID == id {
				fmt.Print("eh esse")

				req, err := io.ReadAll(r.Body)
				if err != nil {
					panic(err)
				}
				defer r.Body.Close()

				var userUpdated User
				err = json.Unmarshal(req, &userUpdated)
				if err != nil {
					panic(err)
				}
				newUser := userUpdated
				users[i].Name = newUser.Name
				fmt.Println("USER ATUALIZADO: ", users[i])
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"message":"user updated!"}`))
				return
			}
		}

		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`Metodo nao encontrado`))

}

func handlerDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		query := r.URL.Query()
		param := query.Get("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			panic(err)
		}

		for i := 0; i < len(users); i++ {
			if users[i].ID == id {
				users = append(users[:i], users[i+1:]...)

			}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`Usuario deletado!`))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`Metodo nao encontrado`))

}

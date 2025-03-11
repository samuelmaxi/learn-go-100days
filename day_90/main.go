package main

import (
	"errors"
	"fmt"
)

type NotFoundError struct {
	Resource string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s nao encotrado", e.Resource)
}

func GetUser(id int) (string, error) {
	if id != 1 {
		return "", &NotFoundError{Resource: "Usuario"}
	}
	return "Usuario valido", nil
}

func main() {
	user, err := GetUser(2)
	if err != nil {
		var notFoundErr *NotFoundError
		if errors.As(err, &notFoundErr) {
			fmt.Println("Erro de recurso não encontrado:", err)
		} else {
			fmt.Println("Erro desconhecido:", err)
		}
		return
	}

	fmt.Println("Usuario encontrado: ", user)

}

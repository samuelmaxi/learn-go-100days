package main

import "fmt"

type User struct {
	Name  string
	Idade int
}

func main() {
	var name string
	var idade int

	fmt.Println("Digite seu nome para cadastro:")
	fmt.Scan(&name)
	fmt.Println("Digite sua idade para cadastro:")
	fmt.Scan(&idade)

	user := User{
		Name:  name,
		Idade: idade,
	}

	fmt.Println("Usuario registrado: ")
	fmt.Println("Nome: ", user.Name)
	fmt.Println("Idade: ", user.Idade)

}

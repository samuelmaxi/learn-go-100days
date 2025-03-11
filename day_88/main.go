package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
)

type User struct {
	Name  string `validate:"required,min=3,max=50"`
	Email string `validate:"required,email"`
	Age   int    `validate:"gte=18,lte=100"`
}

func sanitizeInput(input string) string {
	p := bluemonday.StrictPolicy()
	return strings.TrimSpace(p.Sanitize(input))
}

func validateAlphaSpaces(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[a-zA-ZÀ-ÿ\s]+$`)
	return re.MatchString(fl.Field().String())
}

func main() {
	validate := validator.New()
	validate.RegisterValidation("alpha_spaces", validateAlphaSpaces)

	// TENTATIVA DE INVASAO
	user := User{
		Name:  "Robert'); DROP TABLE users; --",
		Email: "<script>alert('Hacked');</script>@example.com",
		Age:   17,
	}

	// LOGIN NORMAL
	// user := User{
	// 	Name:  "Robert",
	// 	Email: "sla@example.com",
	// 	Age:   18,
	// }

	user.Name = sanitizeInput(user.Name)
	user.Email = sanitizeInput(user.Email)

	err := validate.Struct(user)
	if err != nil {
		fmt.Println("Erro de valdação: ", err)
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("- Campo '%s': %s\n", err.Field(), err.Tag())
		}
	} else {
		fmt.Println("Usuario valido: ", user)
	}

}

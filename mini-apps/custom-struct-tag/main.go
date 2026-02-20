package main

import (
	"fmt"
	"net/mail"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type User struct {
	Name  string `validate:"required,min=2,max=90"`
	Email string `validate:"required,email"`
}

// faghat yee bar complie mikonim baraye behineh boodan
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func isValidEmailByStdLib(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validate(val any) []error {
	v := reflect.ValueOf(val)
	var validationErrs []error

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := v.Type().Field(i).Tag.Get("validate")

		if tag == "" {
			continue
		}

		rules := strings.Split(tag, ",")
		for _, rule := range rules {
			fieldName := v.Type().Field(i).Name

			switch {
			case strings.HasPrefix(rule, "min="):
				min, _ := strconv.Atoi(strings.TrimPrefix(rule, "min="))
				if len(field.String()) < min {
					validationErrs = append(
						validationErrs,
						fmt.Errorf("%s bayad bozorg tar az %d bashe.", fieldName, min),
					)
				}

			case strings.HasPrefix(rule, "max="):
				max, _ := strconv.Atoi(strings.TrimPrefix(rule, "max="))
				if len(field.String()) > max {
					validationErrs = append(
						validationErrs,
						fmt.Errorf("%s bayad cochik tar az %d bashe.", fieldName, max),
					)
				}

			case rule == "required":
				// || field.IsNil() || field.IsZero() || field.IsValid()
				if field.String() == "" {
					validationErrs = append(
						validationErrs,
						fmt.Errorf("%s is required hast hatman pooresh bookon.", fieldName),
					)
				}

			case rule == "email":
				if !isValidEmail(field.String()) {
					validationErrs = append(
						validationErrs,
						fmt.Errorf("%s is not format email is invalid.", fieldName),
					)
				}

			}
		}
	}

	return validationErrs
}

func main() {
	user := User{
		Name:  "Pouya",
		Email: "pouya@mail.com",
	}

	invalidUser := User{
		Name:  "abc",
		Email: "abc@sf.",
	}

	if err := validate(user); err != nil {
		fmt.Printf("ValidationError: %+v\n", err)
	}
	if err := validate(invalidUser); err != nil {
		fmt.Printf("ValidationError: %+v\n", err)
	}
}

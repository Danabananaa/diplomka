package model

import (
	"errors"
	"regexp"
	"unicode"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Validate() error {
	if !validateName(u.Name) {
		return errors.New("invalid name")
	}
	if !validateSurname(u.Surname) {
		return errors.New("invalid surname")
	}
	if !validateEmail(u.Email) {
		return errors.New("invalid email")
	}
	if !validatePassword(u.Password) {
		return errors.New("invalid password")
	}
	return nil
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *Authentication) Validate() error {
	if !validateEmail(a.Email) {
		return errors.New("invalid email")
	}
	if !validatePassword(a.Password) {
		return errors.New("invalid password")
	}
	return nil
}

func validateName(name string) bool {
	regex := regexp.MustCompile(`^\p{Lu}\p{L}*$`)
	return regex.MatchString(name)
}

func validateSurname(surname string) bool {
	regex := regexp.MustCompile(`^\p{Lu}\p{L}*(-\p{Lu}\p{L}*)?$`)
	return regex.MatchString(surname)
}

func validateEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

func validatePassword(password string) bool {
	// Проверяем длину пароля (минимальная длина - 8 символов)
	if len(password) < 8 {
		return false
	}

	// Проверяем, что пароль содержит хотя бы одну букву в верхнем регистре и одну букву в нижнем регистре
	var hasUpper, hasLower bool
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		}
		if hasUpper && hasLower {
			break
		}
	}
	if !hasUpper || !hasLower {
		return false
	}

	// Проверяем, что пароль содержит хотя бы одну цифру
	hasDigit := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
			break
		}
	}

	return hasDigit
}

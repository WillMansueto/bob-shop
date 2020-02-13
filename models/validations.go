package models

import(
	"errors"
	"github.com/badoux/checkmail"
)

var(
	ErrRequiredFirstName = errors.New("Nome requerido")
	ErrRequiredLastName = errors.New("Sobrenome requerido")
	ErrRequiredEmail = errors.New("Email requerido")
	ErrInvalidEmail = errors.New("Email inválido")
	ErrRequiredPassword = errors.New("Senha requerida")
)

func IsEmpty(attr string) bool {
	if attr == "" {
		return true
	}
	return false
}

func IsEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return false
	}
	return true
}

func ValidateNewUser(user User) (User, error) {
	if IsEmpty(user.FirstName) {
		return User{}, ErrRequiredFirstName
	}
	if IsEmpty(user.LastName) {
		return User{}, ErrRequiredLastName
	}
	if IsEmpty(user.Email) {
		return User{}, ErrRequiredEmail
	}
	if !IsEmail(user.Email) {
		return User{}, ErrInvalidEmail
	}
	if IsEmpty(user.Password) {
		return User{}, ErrRequiredPassword
	}
	return user, nil
}
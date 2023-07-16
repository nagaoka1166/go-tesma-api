package entity

import "github.com/google/uuid"

type User struct {
	ID            uuid.UUID `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	FirstNameKana string    `json:"firstNameKana"`
	LastNameKana  string    `json:"lastNameKana"`
	Email         string    `json:"email"`
	BirthDate     string    `json:"birthDate"`
	Faculty       Faculty   `json:"faculty"`
	EmailVerified bool      `json:"emailVerified"`
	SignUpDate    string    `json:"signUpDate"`
}
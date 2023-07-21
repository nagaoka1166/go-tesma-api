package entity

import (
    "github.com/google/uuid"
    "github.com/go-playground/validator/v10"
)

type User struct {
    ID            uuid.UUID `json:"id"`
    FirstName     string    `json:"firstName"`
    LastName      string    `json:"lastName"`
    FirstNameKana string    `json:"firstNameKana"`
    LastNameKana  string    `json:"lastNameKana"`
    Email         string    `json:"email" validate:"required,email,endswith=@ed.ritsumei.ac.jp"`
    Password      string    `json:"password" validate:"required,min=8,alphanum,containsany=0123456789,containsany=abcdefghijklmnopqrstuvwxyz"`
    BirthDate     string    `json:"birthDate"`
    Faculty       Faculty   `json:"faculty"`
    EmailVerified bool      `json:"emailVerified"`
    SignUpDate    string    `json:"signUpDate"`
}

var validate *validator.Validate

func init() {
    validate = validator.New()
}

func (u *User) Validate() error {
    return validate.Struct(u)
}

//  app/domain/entity/user.go
package entity

import (
    "github.com/google/uuid"
    "github.com/go-playground/validator/v10"
    "database/sql"
    "gorm.io/gorm"
)

type User struct {
    ID            uuid.UUID `gorm:"type:char(36);primary_key"`
    FirstName     string    `json:"firstName" gorm:"type:varchar(100)"`
    LastName      string    `json:"lastName" gorm:"type:varchar(100)"`
    FirstNameKana string    `json:"firstNameKana" gorm:"type:varchar(100)"`
    LastNameKana  string    `json:"lastNameKana" gorm:"type:varchar(100)"`
    Email         string    `json:"email" gorm:"type:varchar(255);uniqueIndex" validate:"required,email,endswith=@ed.ritsumei.ac.jp"`
    Password      string    `json:"password" gorm:"type:varchar(255)" validate:"required,min=8,alphanum"`
    BirthDate     sql.NullTime    `json:"birthDate" gorm:"type:date"`
    FacultyID     *int     `json:"facultyID" gorm:"type:int;index"`
    Faculty       *Faculty `gorm:"foreignKey:FacultyID;references:ID"`
    EmailVerified bool      `json:"emailVerified" gorm:"type:boolean;default:false"`
}



func (u *User) GetFaculty() *Faculty {
    return u.Faculty
}

var validate *validator.Validate

func init() {
    validate = validator.New()
}

func (u *User) Validate() error {
    return validate.Struct(u)
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    u.ID = uuid.New()
    return
}
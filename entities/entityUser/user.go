package entityuser

import (
	"time"

	"github.com/faqihyugos/pengaduan-api/entities"
	"github.com/go-playground/validator/v10"
)

type User struct {
	entities.GormModel
	Username string `json:"username" gorm:"unique;not null" `
	Fullname string `json:"full_name" gorm:"not null" `
	Email    string `json:"email" gorm:"unique;not null" `
	Password string `json:"password" gorm:"not null" `
	Nik      string `json:"nik" gorm:"unique;not null" `
	Role     string `json:"role" gorm:"not null" `
	Phone    string `json:"phone" gorm:"not null" `
}

type Request struct {
	ID       uint   `json:"id,omitempty" swaggerignore:"true"`
	Fullname string `json:"full_name" example:"Jhon Doe" form:"full_name" validate:"required,min=3,max=50"`
	Username string `json:"username" example:"jhondoe" form:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" example:"test@example.com" form:"email" validate:"required,email"`
	Password string `json:"password" example:"password" form:"password" validate:"required,min=6,max=50"`
	Nik      string `json:"nik" example:"1234567890123456" form:"nik" validate:"required,min=16,max=16"`
	Phone    string `json:"phone" example:"081234567890" form:"phone" validate:"required,min=11,max=13"`
	Role     string `json:"role" example:"user" form:"role" validate:"required,min=4,max=10"`
}

type Response struct {
	ID        uint       `json:"id"  example:"1"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" example:"2021-11-03T01:52:41.035Z"`
	Fullname  string     `json:"full_name" example:"Jhon Doe"`
	Username  string     `json:"username"  example:"jhondoe"`
	Email     string     `json:"email" example:"test@example.com"`
	Nik       string     `json:"nik" example:"1234567890123456"`
	Phone     string     `json:"phone" example:"081234567890"`
	Role      string     `json:"role" example:"user"`
}

type RequestLogin struct {
	Email    string `json:"email" example:"test@example.com" form:"email" valid:"required-Your email is required"`
	Password string `json:"password" example:"password" form:"password" valid:"required-Your password is required"`
}

type ResponseLogin struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJxd2Vxd2..."`
}

// ExampleRequestUpdate only for example swaggo docs
type ExampleRequestUpdate struct {
	Username string `json:"username" example:"jhondoe"`
	Email    string `json:"email" example:"test@example.com"`
}

type ExampleResponseDelete struct {
	Message string `json:"message" example:"your account has been successfully deleted"`
}

// validate struct request
var validate *validator.Validate = validator.New()

func (u *Request) Validate() error {
	return validate.Struct(u)
}

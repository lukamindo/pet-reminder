package request

import (
	"github.com/lukamindo/pet-reminder/app/db"
)

// UserRegister request struct to register player
type UserRegister struct {
	Username        string `json:"username" validate:"required,min=5,max=25"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6,max=127"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=127"`
}

// DB converts userRegister request to user DB object
func (ur UserRegister) DB(hashedPwd string) db.User {
	return db.User{
		Username: ur.Username,
		Password: hashedPwd,
		Email:    ur.Email,
	}
}

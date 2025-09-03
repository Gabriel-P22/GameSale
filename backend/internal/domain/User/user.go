package domain_user

import (
	"game-sale-backend/internal/interfaces/controllers/dtos"

	"github.com/google/uuid"
)

type User struct {
	ID string
	Name string
	Email string
	password string
}


func NewUser(name, email, passowrd string) *User {
	return &User{
		ID: uuid.New().String(),
		Name: name,
		Email: email,
		password: passowrd,
	} 
}

func (u *User) ToUserResponse() *dtos.UserResponse {
	return &dtos.UserResponse{
		ID: u.ID,
		Name: u.Name,
		Email: u.Email,
	}
}

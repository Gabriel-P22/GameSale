package application_user

import (
	"game-sale-backend/internal/interfaces/controllers/dtos"
)

type UserCommand struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func CreateCommand(user dtos.UserRequest) *UserCommand {
	return &UserCommand{
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}
}


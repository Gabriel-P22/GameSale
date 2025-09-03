package application_user

import domain_user "game-sale-backend/internal/domain/User"

type UserService struct {}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) Create(command *UserCommand) (*domain_user.User) {
	return domain_user.NewUser(command.Name, command.Email, command.Password);
}

package domain_user

import (
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
package request

import "go-crud/internal/model"

type User struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func FromUserModel(input model.User) User {
	return User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
}

func ToUserModel(input User) model.User {
	return model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
}

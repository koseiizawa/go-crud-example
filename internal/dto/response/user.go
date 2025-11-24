package response

import "go-crud/internal/model"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromUserModel(input model.User) User {
	return User{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
	}
}

func ToUserModel(input User) model.User {
	return model.User{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
	}
}

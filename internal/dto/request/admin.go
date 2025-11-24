package request

import "go-crud/internal/model"

type Admin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func FromAdminModel(input model.Admin) Admin {
	return Admin{
		Email:    input.Email,
		Password: input.Password,
	}
}

func ToAdminModel(input Admin) model.Admin {
	return model.Admin{
		Email:    input.Email,
		Password: input.Password,
	}
}

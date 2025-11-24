package response

import "go-crud/internal/model"

type Admin struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func FromAdminModel(input model.Admin) Admin {
	return Admin{
		ID:    input.ID,
		Email: input.Email,
	}
}

func ToAdminModel(input Admin) model.Admin {
	return model.Admin{
		ID:    input.ID,
		Email: input.Email,
	}
}

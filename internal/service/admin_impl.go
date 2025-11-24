package service

import (
	"go-crud/internal/dto/request"
	"go-crud/internal/dto/response"
	"go-crud/internal/repository"
)

type AdminServiceImpl struct {
	Repository repository.AdminRepository
}

func (s *AdminServiceImpl) SignUp(admin request.Admin) (response.Admin, error) {
	modelUser := request.ToAdminModel(admin)
	createdUser, err := s.Repository.SignUp(modelUser)
	if err != nil {
		return response.Admin{}, err
	}

	return response.FromAdminModel(createdUser), nil
}

func (s *AdminServiceImpl) LogIn(admin request.Admin) (response.Admin, error) {
	modelUser := request.ToAdminModel(admin)
	existingUser, err := s.Repository.LogIn(modelUser)
	if err != nil {
		return response.Admin{}, err
	}

	return response.FromAdminModel(existingUser), nil
}

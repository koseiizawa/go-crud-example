package service

import (
	"go-crud/internal/domain"
	"go-crud/internal/dto"
	"go-crud/internal/repository"
)

type AdminServiceImpl struct {
	Repository repository.AdminRepository
}

func (s *AdminServiceImpl) SignUp(user dto.AdminRequest) (dto.AdminResponse, error) {
	domainUser := AdminDtoRequestToDomain(user)
	createdUser, err := s.Repository.SignUp(domainUser)
	if err != nil {
		return dto.AdminResponse{}, err
	}

	return AdminDomainToDtoResponse(createdUser), nil
}

func (s *AdminServiceImpl) LogIn(user dto.AdminRequest) (dto.AdminResponse, error) {
	domainUser := AdminDtoRequestToDomain(user)
	existingUser, err := s.Repository.LogIn(domainUser)
	if err != nil {
		return dto.AdminResponse{}, err
	}

	return AdminDomainToDtoResponse(existingUser), nil
}

func AdminDtoRequestToDomain(input dto.AdminRequest) domain.AdminUser {
	return domain.AdminUser{
		Email:    input.Email,
		Password: input.Password,
	}
}

func AdminDomainToDtoResponse(input domain.AdminUser) dto.AdminResponse {
	return dto.AdminResponse{
		ID:    input.ID,
		Email: input.Email,
	}
}

package service

import (
	"go-crud/internal/domain"
	"go-crud/internal/dto"
	"go-crud/internal/repository"
)

type UserServiceImpl struct {
	Repository repository.UserRepository
}

func (s *UserServiceImpl) Create(user dto.UserRequest) (dto.UserResponse, error) {
	createdUser, err := s.Repository.Create(toDomainStruct(user))
	return toResponseStruct(createdUser), err
}

func (s *UserServiceImpl) GetAll() ([]dto.UserResponse, error) {
	users, err := s.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	responseUsers := make([]dto.UserResponse, len(users))
	for i, v := range users {
		responseUsers[i] = toResponseStruct(v)
	}

	return responseUsers, nil
}

func (s *UserServiceImpl) GetById(id string) (dto.UserResponse, error) {
	user, err := s.Repository.GetById(id)
	return toResponseStruct(user), err
}

func (s *UserServiceImpl) Update(user dto.UserRequest, id string) (dto.UserResponse, error) {
	updatedUser, err := s.Repository.Update(toDomainStruct(user), id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return toResponseStruct(updatedUser), nil
}

func (s *UserServiceImpl) Delete(id string) error {
	return s.Repository.Delete(id)
}

func toDomainStruct(input dto.UserRequest) domain.User {
	return domain.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
}

func toResponseStruct(input domain.User) dto.UserResponse {
	return dto.UserResponse{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
	}
}

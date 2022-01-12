/*
User Service. Application layer
*/
package user

import "microservice-template-ddd/internal/user/domain"

type Service struct{}

func New() (*Service, error) {
	return &Service{}, nil
}

func (s *Service) Get() (*domain.User, error) {
	return &domain.User{
		Login:    "test@user",
		Password: "",
		Email:    "test@user.com",
		IsActive: true,
	}, nil
}

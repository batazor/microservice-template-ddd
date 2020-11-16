/*
User Service. Application layer
*/
package application

import "robovoice-template/internal/user/domain"

type Service struct{}

func (s *Service) Get() (*domain.User, error) {
	return &domain.User{
		Login:    "test@user",
		Password: "",
		Email:    "test@user.com",
		IsActive: true,
	}, nil
}

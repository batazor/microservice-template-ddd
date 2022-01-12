/*
Billing Service. Application layer
*/
package billing

import "microservice-template-ddd/internal/billing/domain"

type Service struct{}

func New() (*Service, error) {
	return &Service{}, nil
}

func (s *Service) Get() (*domain.Billing, error) {
	return &domain.Billing{
		Balance: 100.00,
	}, nil
}

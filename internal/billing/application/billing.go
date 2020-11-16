/*
Billing Service. Application layer
*/
package application

import "robovoice-template/internal/billing/domain"

type Service struct{}

func (s *Service) Get() (*domain.Billing, error) {
	return &domain.Billing{
		Balance: 100.00,
	}, nil
}

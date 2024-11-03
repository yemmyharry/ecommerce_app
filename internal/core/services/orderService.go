package services

import (
	mysql_repo "ecommerce/internal/adapter/repositories/mysql"
	"ecommerce/internal/core/domain"
	"ecommerce/internal/ports"
)

type OrderService struct {
	repo ports.OrderRepository
}

func NewOrderService() *OrderService {
	return &OrderService{
		repo: mysql_repo.NewOrderRepository(),
	}
}

func (s *OrderService) GetAll(params map[string]interface{}) ([]domain.Order, error) {
	return s.repo.GetAll(params)
}

func (s *OrderService) Find(id string) (*domain.Order, error) {
	return s.repo.Find(id)
}

func (s *OrderService) Create(input *domain.Order) (*domain.Order, error) {
	return s.repo.Create(input)
}

func (s *OrderService) Update(id string, input *domain.Order) (*domain.Order, error) {
	return s.repo.Update(id, input)
}

func (s *OrderService) Delete(id string) error {
	return s.repo.Delete(id)
}

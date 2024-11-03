package services

import (
	"ecommerce/internal/core/domain"
	mysql_repo "ecommerce/internal/adapter/repositories/mysql"
	"ecommerce/internal/ports"
)

type ProductService struct {
	repo ports.ProductRepository
}

func NewProductService() *ProductService {
	return &ProductService{
		repo: mysql_repo.NewProductRepository(),
	}
}

func (s *ProductService) GetAll(params map[string]interface{}) ([]domain.Product, error) {
	return s.repo.GetAll(params)
}

func (s *ProductService) Find(id string) (*domain.Product, error) {
	return s.repo.Find(id)
}

func (s *ProductService) Create(input *domain.Product) (*domain.Product, error) {
	return s.repo.Create(input)
}

func (s *ProductService) Update(id string, input *domain.Product) (*domain.Product, error) {
	return s.repo.Update(id, input)
}

func (s *ProductService) Delete(id string) error {
	return s.repo.Delete(id)
}
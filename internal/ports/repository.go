package ports

import "ecommerce/internal/core/domain"

type UserRepository interface {
	Create(data interface{}) (interface{}, error)
	Find(id string) (interface{}, error)
	GetAll(param map[string]interface{}) (interface{}, error)
	Update(id string, data interface{}) (interface{}, error)
	Delete(id string) (interface{}, error)
}

type ProductRepository interface {
	Create(data *domain.Product) (*domain.Product, error)
	Find(id string) (*domain.Product, error)
	GetAll(param map[string]interface{}) ([]domain.Product, error)
	Update(id string, data *domain.Product) (*domain.Product, error)
	Delete(id string) error
}

type OrderRepository interface {
	Create(data *domain.Order) (*domain.Order, error)
	Find(id string) (*domain.Order, error)
	GetAll(param map[string]interface{}) ([]domain.Order, error)
	Update(id string, data *domain.Order) (*domain.Order, error)
	Delete(id string) error
}

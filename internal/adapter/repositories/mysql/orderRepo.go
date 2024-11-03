package mysql_repo

import (
	"ecommerce/internal/core/domain"
	"ecommerce/internal/ports"
	"gorm.io/gorm/clause"
)

type OrderRepository struct{}

func NewOrderRepository() ports.OrderRepository {
	return &OrderRepository{}
}

func (repo *OrderRepository) GetAll(params map[string]interface{}) ([]domain.Order, error) {
	var orders []domain.Order
	q := DB.Preload(clause.Associations)

	if userID, ok := params["user_id"]; ok {
		q = q.Where("user_id = ?", userID)
	}

	if err := q.Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (repo *OrderRepository) Find(id string) (*domain.Order, error) {
	var order domain.Order
	if err := DB.Preload(clause.Associations).Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (repo *OrderRepository) Create(data *domain.Order) (*domain.Order, error) {
	if err := DB.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *OrderRepository) Update(id string, data *domain.Order) (*domain.Order, error) {
	var order domain.Order
	if err := DB.Model(&order).Where("id = ?", id).Updates(data).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (repo *OrderRepository) Delete(id string) error {
	if err := DB.Where("id = ?", id).Delete(&domain.Order{}).Error; err != nil {
		return err
	}
	return nil
}

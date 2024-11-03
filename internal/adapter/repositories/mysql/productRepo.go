package mysql_repo

import (
	"ecommerce/internal/core/domain"
	"ecommerce/internal/ports"
	"gorm.io/gorm/clause"
)

type ProductRepository struct{}

func NewProductRepository() ports.ProductRepository {
	return &ProductRepository{}
}

func (repo *ProductRepository) GetAll(params map[string]interface{}) ([]domain.Product, error) {
	var products []domain.Product
	q := DB.Preload(clause.Associations)

	if name, ok := params["name"]; ok {
		q = q.Where("name = ?", name)
	}

	if err := q.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (repo *ProductRepository) Find(id string) (*domain.Product, error) {
	var product domain.Product
	if err := DB.Preload(clause.Associations).Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo *ProductRepository) Create(data *domain.Product) (*domain.Product, error) {
	if err := DB.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *ProductRepository) Update(id string, data *domain.Product) (*domain.Product, error) {
	var product domain.Product
	if err := DB.Model(&product).Where("id = ?", id).Updates(data).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo *ProductRepository) Delete(id string) error {
	if err := DB.Where("id = ?", id).Delete(&domain.Product{}).Error; err != nil {
		return err
	}
	return nil
}

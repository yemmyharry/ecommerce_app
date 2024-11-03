package mysql_repo

import (
	"ecommerce/internal/core/domain"
	"ecommerce/internal/ports"
	"fmt"

	"gorm.io/gorm/clause"
)

type UserRepository struct {
}

func NewUserRepository() ports.UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) GetAll(param map[string]interface{}) (interface{}, error) {
	model := repo.ArrayModel()

	fmt.Println(DB)

	q := DB.Preload(clause.Associations)

	if param["name"] != nil {
		q.Where("name = ?", param["name"])
	}

	if param["phone"] != nil {
		q.Where("phone = ?", param["phone"])
	}

	if param["email"] != nil {
		q.Where("email = ?", param["email"])
	}

	q.Find(&model)
	return model, q.Error
}

func (repo *UserRepository) Find(id string) (interface{}, error) {
	model := repo.Model()
	q := DB.Preload(clause.Associations).Where("id = ?", id).First(&model)

	return model, q.Error
}

func (repo *UserRepository) Create(data interface{}) (interface{}, error) {
	model := repo.Model()
	q := DB.Model(model).Create(data)
	return data, q.Error

}

func (repo *UserRepository) Update(id string, data interface{}) (interface{}, error) {
	model := repo.Model()
	q := DB.Model(&model).Where("id = ?", id).Updates(data)
	if q.Error != nil {
		return nil, q.Error
	}
	return repo.Find(id)
}

func (repo *UserRepository) Delete(id string) (interface{}, error) {
	model := repo.Model()
	q := DB.Model(&model).Where("id = ?", id).Delete(model)
	return model, q.Error
}

func (repo *UserRepository) Model() domain.User {
	return domain.User{}
}

func (repo *UserRepository) ArrayModel() []domain.User {
	return []domain.User{}
}

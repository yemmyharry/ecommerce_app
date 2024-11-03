package services

import (
	"ecommerce/internal/adapter/api/requests"
	mysql_repo "ecommerce/internal/adapter/repositories/mysql"
	"ecommerce/internal/core/domain"
	"ecommerce/internal/ports"
	"errors"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repo: mysql_repo.NewUserRepository(),
	}
}

func (s *UserService) Save(input domain.User) (interface{}, error) {
	return s.repo.Create(&input)
}

func (s *UserService) Update(id string, input requests.UpdateUserRequest) (interface{}, error) {

	return s.repo.Update(id, input)
}

func (s *UserService) GetAll(param map[string]interface{}) (interface{}, error) {
	return s.repo.GetAll(param)
}

func (s *UserService) Find(id string) (interface{}, error) {
	return s.repo.Find(id)
}

func (s *UserService) FindUserByEmailORPhone(param string) (*domain.User, error) {
	users := []domain.User{}
	data, err := s.GetAll(map[string]interface{}{"phone": param})
	if err != nil {
		data, err = s.GetAll(map[string]interface{}{"email": param})
		if err != nil {
			return nil, errors.New("error fetching users")
		}
	}
	users, ok := data.([]domain.User)
	if !ok {
		return nil, errors.New("error decoding users")
	}

	if len(users) == 0 {
		return nil, errors.New("user not found")
	}

	return &users[0], nil
}

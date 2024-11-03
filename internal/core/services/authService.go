package services

import (
	"ecommerce/internal/adapter/api/requests"
	"ecommerce/internal/core/domain"
	"ecommerce/internal/helper"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userService *UserService
}

func NewAuthService() *AuthService {
	return &AuthService{
		userService: NewUserService(),
	}
}

func (s *AuthService) Register(input requests.SignUpRequest) (interface{}, error) {

	_, err := s.userService.FindUserByEmailORPhone(input.Phone)
	if err == nil {
		return nil, errors.New("user already exits")
	}
	fmt.Println(input.Phone)
	user := domain.User{
		Email:     &input.Email,
		Lastname:  input.Lastname,
		Firstname: input.Firstname,
		Phone:     input.Phone,
		IsAdmin:   input.IsAdmin,
	}

	//Generate and hash password
	b, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	if err != nil {
		return nil, errors.New("an error occurred hashing password")
	}
	user.Password = string(b)

	data, err := s.userService.Save(user)
	if err != nil {
		return nil, errors.New("error occurred creating user")
	}

	token, err := helper.GenerateJWT(user)
	if err != nil {
		return nil, errors.New("error generating token")
	}

	return map[string]interface{}{
		"data":  data,
		"token": token,
	}, err

}

func (s *AuthService) Login(input requests.LoginRequest) (interface{}, error) {
	fmt.Println(input.Phone)
	user, err := s.userService.FindUserByEmailORPhone(input.Phone)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	token, err := helper.GenerateJWT(user)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"data":  user,
		"token": token,
	}, nil
}

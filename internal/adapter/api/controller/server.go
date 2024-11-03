package controller

import "ecommerce/internal/core/services"

type HTTPHandler struct {
	userHandler    *UserController
	orderHandler   *OrderController
	productHandler *ProductController
	auth           *services.AuthService
}

func NewHTTPHandler(options ...interface{}) *HTTPHandler {
	handler := &HTTPHandler{
		userHandler:    NewUserController(),
		auth:           services.NewAuthService(),
		orderHandler:   NewOrderController(),
		productHandler: NewProductController(),
	}
	return handler
}

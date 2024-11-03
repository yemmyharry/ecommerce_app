package controller

import (
	"ecommerce/internal/adapter/api/middleware"

	"github.com/gin-gonic/gin"
)

func (s *HTTPHandler) Routes(router *gin.Engine) {
	router.Use(middleware.LoggingMiddleWare())
	api := router.Group("/api/v1")

	api.GET("/healthcheck", s.HealthCheck)
	api.POST("/register", s.Register)
	api.POST("/login", s.Login)

	//user routers
	users := api.Group("/users")
	{
		users.Use(middleware.AuthMiddleware())
		users.GET("/", s.userHandler.GetAllUsers)
		users.PUT("/:id", s.userHandler.UpdateUser)
		users.GET("/:id", s.userHandler.FindUser)
	}

	// Product routes
	products := api.Group("/products")
	{
		// Admin-only product routes
		products.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		{
			products.POST("/", s.productHandler.CreateProduct)
			products.PUT("/:id", s.productHandler.UpdateProduct)
			products.GET("/", s.productHandler.GetAllProducts)
			products.GET("/:id", s.productHandler.FindProduct)
			products.DELETE("/:id", s.productHandler.DeleteProduct)
		}
	}

	// Order routes
	orders := api.Group("/orders")
	orders.Use(middleware.AuthMiddleware())
	{
		orders.POST("/", s.orderHandler.CreateOrder)
		orders.GET("/", s.orderHandler.GetAllOrders)
		orders.GET("/:id", s.orderHandler.FindOrder)
		orders.DELETE("/:id", s.orderHandler.DeleteOrder)

		// Admin-only order management
		orders.Use(middleware.AdminMiddleware())
		{
			orders.PUT("/:id/status", s.orderHandler.UpdateOrderStatus) // Admin-only: Update the status of an order
		}
	}

	router.NoRoute(func(c *gin.Context) { c.String(404, "Not found") })
}

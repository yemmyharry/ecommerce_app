package controller

import (
	"ecommerce/internal/core/domain"
	"ecommerce/internal/core/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController struct {
	service *services.OrderService
}

func NewOrderController() *OrderController {
	return &OrderController{
		service: services.NewOrderService(),
	}
}

func (o *OrderController) GetAllOrders(c *gin.Context) {
	params := FlatUrlQuery(c.Request.URL.Query())
	resp, err := o.service.GetAll(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (o *OrderController) FindOrder(c *gin.Context) {
	id := c.Param("id")
	resp, err := o.service.Find(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (o *OrderController) CreateOrder(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := o.service.Create(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (o *OrderController) UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := o.service.Update(id, &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (o *OrderController) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := o.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}

func (o *OrderController) UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedOrder, err := o.service.Update(id, &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedOrder)
}

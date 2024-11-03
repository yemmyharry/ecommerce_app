package controller

import (
	"ecommerce/internal/adapter/api/requests"
	"ecommerce/internal/adapter/api/response"
	"ecommerce/internal/core/services"
	"ecommerce/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		service: services.NewUserService(),
	}
}

func (s *UserController) GetAllUsers(c *gin.Context) {
	logger.Info("get all users")

	params := FlatUrlQuery(c.Request.URL.Query())

	resp, err := s.service.GetAll(params)
	if err != nil {
		logger.Error("Error saving sprint" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewUserArrayResponse(resp, err))
}

func (s *UserController) UpdateUser(c *gin.Context) {
	logger.Info("update user")

	request := requests.UpdateUserRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	params := c.Param("id")

	resp, err := s.service.Update(params, request)
	if err != nil {
		logger.Error("Error saving sprint" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewUserResponse(resp, err))
}

func (s *UserController) FindUser(c *gin.Context) {
	logger.Info("find user")

	id := c.Param("id")

	resp, err := s.service.Find(id)
	if err != nil {
		logger.Error("Error saving sprint" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewUserResponse(resp, err))
}

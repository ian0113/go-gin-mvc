package controllers

import (
	"net/http"
	"strconv"

	"github.com/ian0113/go-gin-mvc/infra"
	"github.com/ian0113/go-gin-mvc/services"
	"github.com/ian0113/go-gin-mvc/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	logger      *zap.Logger
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		logger:      infra.GetLogger().Named("user.controler"),
		userService: services.NewUserService(),
	}
}

// POST /api/users
func (x *UserController) Register(c *gin.Context) {
	type UserRegisterRequest struct {
		Name     string `json:"name" binding:"required,min=6"`
		Email    string `json:"email" binding:"required,email"`
		Account  string `json:"account" binding:"required,min=6"`
		Password string `json:"password" binding:"required,min=6"`
	}
	req := UserRegisterRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := x.userService.CreateUser(req.Name, req.Email, req.Account, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":      user.ID,
		"name":    user.Name,
		"email":   user.Email,
		"account": user.Account,
	})
}

// DELETE /api/users/:id
func (x *UserController) Unregister(c *gin.Context) {
	reqIDStr, ok := c.Params.Get("id")
	reqID, err := strconv.ParseUint(reqIDStr, 10, 32)
	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userID, ok := utils.GetGinContextUserID(c)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user id"})
		return
	}

	if reqID != uint64(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid request"})
		return
	}

	err = x.userService.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

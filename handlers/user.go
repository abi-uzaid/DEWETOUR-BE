package handlers

import (
	resultdto "dewetour/dto/result"
	usersdto "dewetour/dto/users"
	"dewetour/models"
	"dewetour/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserRepository repositories.UserRepository
}

func NewUserHandler(userRepository repositories.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepository: userRepository,
	}
}

func (h *UserHandler) FindUsers(c *gin.Context) {
	users, err := h.UserRepository.FindUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: users})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")

	ID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.UserRepository.GetUser(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: user})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	ID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.UserRepository.GetUser(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.UserRepository.DeleteUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponse(user)})
}

func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:       u.ID,
		Fullname: u.Fullname,
		Email:    u.Email,
		Password: u.Password,
		Phone:    u.Phone,
		Address:  u.Address,
	}
}

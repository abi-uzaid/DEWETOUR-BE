package handlers

import (
	authdto "dewetour/dto/auth"
	dto "dewetour/dto/result"
	"dewetour/models"
	"dewetour/pkg/bcrypt"
	jwtToken "dewetour/pkg/jwt"
	"dewetour/repositories"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(c *gin.Context) {
	var request authdto.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validation := validator.New()
	if err := validation.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := bcrypt.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Fullname: request.Fullname,
		Email:    request.Email,
		Password: password,
		Phone:    request.Phone,
		Address:  request.Address,
		Role:     request.Role,
	}

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dto.SuccessResult{Code: http.StatusOK, Data: data.ID}})
}

func (h *handlerAuth) Login(c *gin.Context) {
	var request authdto.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
		return
	}

	// Check Pass
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password wrong"})
		return
	}

	//Generate Token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 24 jam expired

	token, errGenerateToken := jwtToken.GenerateToken(claims)
	if errGenerateToken != nil {
		c.Error(errGenerateToken)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	loginResponse := authdto.LoginResponse{
		Email: user.Email,
		Token: token,
		Role:  user.Role,
	}

	c.JSON(http.StatusOK, gin.H{"data": dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}})
}

func (h *handlerAuth) CheckAuth(c *gin.Context) {
	userInfo, _ := c.Get("userInfo")
	userId := int(userInfo.(jwt.MapClaims)["id"].(float64))

	// Check User by Id
	user, err := h.AuthRepository.CheckAuth(int(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	CheckAuthResponse := authdto.CheckAuthResponse{
		ID:       user.ID,
		Fullname: user.Fullname,
		Email:    user.Email,
		Role:     user.Role,
	}

	c.JSON(http.StatusOK, gin.H{"data": dto.SuccessResult{Code: http.StatusOK, Data: CheckAuthResponse}})
}

// func (h *handlerAuth) LoginAdmin(c *gin.Context) {
// 	var request authdto.LoginRequest

// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user := models.User{
// 		Email:    request.Email,
// 		Password: request.Password,
// 		Role:     "admin",
// 	}

// 	user, err := h.AuthRepository.Login(user.Email)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
// 		return
// 	}

// 	// Check Pass
// 	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
// 	if !isValid {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Password wrong"})
// 		return
// 	}

// 	//Generate Token
// 	claims := jwt.MapClaims{}
// 	claims["id"] = user.ID
// 	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 24 jam expired

// 	token, errGenerateToken := jwtToken.GenerateToken(claims)
// 	if errGenerateToken != nil {
// 		c.Error(errGenerateToken)
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 		return
// 	}

// 	loginResponse := authdto.LoginResponse{
// 		Email: user.Email,
// 		Token: token,
// 		Role:  "admin",
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}})
// }

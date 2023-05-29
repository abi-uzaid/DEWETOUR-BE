package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/middleware"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	userRepository := repositories.NewAuthRepository(mysql.DB)
	Handlers := handlers.HandlerAuth(userRepository)

	r.POST("/register", Handlers.Register)
	r.POST("/login", Handlers.Login)
	r.POST("/login-admin", Handlers.LoginAdmin)
	r.GET("/check-auth", middleware.Auth(Handlers.CheckAuth))
}

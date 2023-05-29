package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/middleware"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gin-gonic/gin"
)

func TransacRoute(r *gin.RouterGroup) {
	TransacRepository := repositories.NewTransactionRepository(mysql.DB)
	Handlers := handlers.NewHandleTransac(TransacRepository)

	r.GET("/transactions", Handlers.FindTransaction)
	r.GET("/transaction/:id", Handlers.GetTransaction)
	r.POST("/transaction", middleware.UploadFile(Handlers.CreateTransaction))
	r.DELETE("/transaction/:id", middleware.Auth(Handlers.DeleteTransaction))
	// r.HandleFunc("/transaction/{id}", Handlers.UpdateTransaction).Methods("PATCH")
	// r.POST("/notification", Handlers.Notification)
}

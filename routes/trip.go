package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/middleware"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gin-gonic/gin"
)

func TripRoute(r *gin.RouterGroup) {

	tripRepository := repositories.NewTripRepository(mysql.DB)
	Handlers := handlers.HandleTrip(tripRepository)

	r.GET("/trips", Handlers.FindTrip)
	r.GET("/trip/:id", Handlers.GetTrip)
	r.POST("/trip", middleware.UploadFile(Handlers.CreateTrip))
	r.PATCH("/trip/:id", Handlers.UpdateTrip)
	r.DELETE("/trip/:id", Handlers.DeleteTrip)
}

package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/middleware"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gin-gonic/gin"
)

func CountryRoutes(r *gin.RouterGroup) {
	CountryReposito := repositories.NewCountryRepository(mysql.DB)
	Handlers := handlers.HandlerCountry(CountryReposito)

	r.POST("/country", middleware.Auth(Handlers.CreateCountry))
	r.GET("/countries", middleware.Auth(Handlers.FindCountry))
	r.PATCH("/country/:id", middleware.Auth(Handlers.UpdateCountry))
	r.GET("/country/:id", middleware.Auth(Handlers.GetCountry))
	r.DELETE("/country/:id", middleware.Auth(Handlers.DeleteCountry))
}

package handlers

import (
	countrydto "dewetour/dto/country"
	resultdto "dewetour/dto/result"
	"dewetour/models"
	"dewetour/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handlerCountry struct {
	CountryReposito repositories.CountryReposito
}

func HandlerCountry(CountryReposito repositories.CountryReposito) *handlerCountry {
	return &handlerCountry{CountryReposito}
}

func (h *handlerCountry) FindCountry(c *gin.Context) {
	country, err := h.CountryReposito.FindCountry()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": country})
}

func (h *handlerCountry) CreateCountry(c *gin.Context) {
	var request countrydto.CountryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validation := validator.New()
	if err := validation.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	country := models.Country{
		Name: request.Name,
	}

	data, err := h.CountryReposito.CreateCountry(country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: countryResConv(data)})
	// c.JSON(http.StatusOK, gin.H{"data": resultdto.SuccessResult{Code: http.StatusOK, Data: countryResConv(data)}})
}

func countryResConv(u models.Country) countrydto.CountryResponse {
	return countrydto.CountryResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}

func (h *handlerCountry) GetCountry(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	country, err := h.CountryReposito.GetCountry(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": country})
}

func (h *handlerCountry) UpdateCountry(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(countrydto.UpdateCountryReq)
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	country, err := h.CountryReposito.GetCountry(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Name != "" {
		country.Name = request.Name
	}

	data, err := h.CountryReposito.UpdateCountry(country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": countryResConv(data)})
}

func (h *handlerCountry) DeleteCountry(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	country, err := h.CountryReposito.GetCountry(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := h.CountryReposito.DeleteCountry(country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

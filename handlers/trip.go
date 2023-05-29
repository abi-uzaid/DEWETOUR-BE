package handlers

import (
	resultdto "dewetour/dto/result"
	tripdto "dewetour/dto/trip"
	"dewetour/models"
	"dewetour/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handleTrip struct {
	TripRepository repositories.TripRepository
}

func HandleTrip(TripRepository repositories.TripRepository) *handleTrip {
	return &handleTrip{TripRepository}
}

func (h *handleTrip) FindTrip(c *gin.Context) {
	trip, err := h.TripRepository.FindTrip()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i, p := range trip {
		trip[i].Image = p.Image
	}

	response := resultdto.SuccessResult{Code: http.StatusOK, Data: trip}
	c.JSON(http.StatusOK, response)
}

func (h *handleTrip) GetTrip(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := resultdto.SuccessResult{Code: http.StatusOK, Data: trip}
	c.JSON(http.StatusOK, response)
}

func (h *handleTrip) DeleteTrip(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := h.TripRepository.DeleteTrip(trip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)}
	c.JSON(http.StatusOK, response)
}

func convertResponseTrip(u models.Trip) tripdto.TripResponse {
	return tripdto.TripResponse{
		ID:           u.ID,
		Title:        u.Title,
		CountryId:    u.CountryId,
		Accomodation: u.Accomodation,
		Transport:    u.Transportation,
		Eat:          u.Eat,
		Day:          u.Day,
		Night:        u.Night,
		Date:         u.DateTrip,
		Price:        u.Price,
		Quota:        u.Quota,
		Description:  u.Description,
		Image:        u.Image,
	}
}

func (h *handleTrip) CreateTrip(c *gin.Context) {

	dataFile := c.GetString("dataFile")

	country_id, _ := strconv.Atoi(c.PostForm("country_id"))
	price, _ := strconv.Atoi(c.PostForm("price"))
	quota, _ := strconv.Atoi(c.PostForm("quota"))
	day, _ := strconv.Atoi(c.PostForm("day"))
	night, _ := strconv.Atoi(c.PostForm("night"))

	request := tripdto.TripRequest{
		Title:        c.PostForm("title"),
		CountryId:    country_id,
		Accomodation: c.PostForm("accomodation"),
		Transport:    c.PostForm("transportation"),
		Eat:          c.PostForm("eat"),
		Day:          day,
		Night:        night,
		Date:         c.PostForm("date"),
		Price:        price,
		Quota:        quota,
		Description:  c.PostForm("desc"),
		Image:        dataFile,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	trip := models.Trip{
		Title:          request.Title,
		CountryId:      request.CountryId,
		Accomodation:   request.Accomodation,
		Transportation: request.Transport,
		Eat:            request.Eat,
		Day:            request.Day,
		Night:          request.Night,
		DateTrip:       request.Date,
		Price:          request.Price,
		Quota:          request.Quota,
		Description:    request.Description,
		Image:          request.Image,
	}

	data, err := h.TripRepository.CreateTrip(trip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)}
	c.JSON(http.StatusOK, response)
}

func (h *handleTrip) UpdateTrip(c *gin.Context) {

	dataFile := c.GetString("dataFile")

	id, _ := strconv.Atoi(c.Param("id"))

	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	country_id, _ := strconv.Atoi(c.PostForm("country_id"))
	price, _ := strconv.Atoi(c.PostForm("price"))
	quota, _ := strconv.Atoi(c.PostForm("quota"))
	day, _ := strconv.Atoi(c.PostForm("day"))
	night, _ := strconv.Atoi(c.PostForm("night"))

	request := tripdto.TripRequest{
		Title:        c.PostForm("title"),
		CountryId:    country_id,
		Accomodation: c.PostForm("accomodation"),
		Transport:    c.PostForm("transportation"),
		Eat:          c.PostForm("eat"),
		Day:          day,
		Night:        night,
		Date:         c.PostForm("date"),
		Price:        price,
		Quota:        quota,
		Description:  c.PostForm("desc"),
		Image:        dataFile,
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Title != "" {
		trip.Title = request.Title
	}

	if request.CountryId != 0 {
		trip.CountryId = request.CountryId
	}

	if request.Accomodation != "" {
		trip.Accomodation = request.Accomodation
	}

	if request.Transport != "" {
		trip.Transportation = request.Transport
	}

	if request.Eat != "" {
		trip.Eat = request.Eat
	}

	if request.Day != 0 {
		trip.Day = request.Day
	}

	if request.Night != 0 {
		trip.Night = request.Night
	}

	if request.Date != "" {
		trip.DateTrip = request.Date
	}

	if request.Price != 0 {
		trip.Price = request.Price
	}

	if request.Quota != 0 {
		trip.Quota = request.Quota
	}

	if request.Description != "" {
		trip.Description = request.Description
	}

	if request.Image != "" {
		trip.Image = request.Image
	}

	data, err := h.TripRepository.UpdateTrip(trip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)}
	c.JSON(http.StatusOK, response)
}

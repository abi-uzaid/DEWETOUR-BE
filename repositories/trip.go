package repositories

import (
	"dewetour/models"

	"gorm.io/gorm"
)

type TripRepository interface {
	FindTrip() ([]models.Trip, error)
	GetTrip(ID int) (models.Trip, error)
	CreateTrip(trip models.Trip) (models.Trip, error)
	UpdateTrip(trip models.Trip) (models.Trip, error)
	DeleteTrip(trip models.Trip) (models.Trip, error)
}

func NewTripRepository(db *gorm.DB) *repository {
	return &repository{db}
}

type repository struct {
	db *gorm.DB
}

func (r *repository) FindTrip() ([]models.Trip, error) {
	var Trip []models.Trip
	err := r.db.Preload("Country").Find(&Trip).Error

	return Trip, err
}

func (r *repository) GetTrip(ID int) (models.Trip, error) {
	var Trip models.Trip
	err := r.db.Preload("Country").First(&Trip, ID).Error

	return Trip, err
}

func (r *repository) CreateTrip(trip models.Trip) (models.Trip, error) {
	err := r.db.Preload("Country").Create(&trip).Error

	return trip, err
}

func (r *repository) DeleteTrip(trip models.Trip) (models.Trip, error) {
	err := r.db.Delete(&trip).Error

	return trip, err
}

func (r *repository) UpdateTrip(trip models.Trip) (models.Trip, error) {
	err := r.db.Save(&trip).Error

	return trip, err
}

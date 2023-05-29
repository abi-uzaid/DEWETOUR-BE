package repositories

import (
	"dewetour/models"

	"gorm.io/gorm"
)

type TransactionRepo interface {
	CreateTransaction(transac models.Transaction) (models.Transaction, error)
	FindTransaction() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	DeleteTransaction(transac models.Transaction) (models.Transaction, error)
	// UpdateTransaction(status string, ID int) (models.Transaction, error)
	GetOneTransaction(ID string) (models.Transaction, error)
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

type TransactionRepository struct {
	db *gorm.DB
}

func (r *TransactionRepository) CreateTransaction(transac models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Trip.Country").Preload("User").Create(&transac).Error

	return transac, err
}

func (r *TransactionRepository) FindTransaction() ([]models.Transaction, error) {
	var Transac []models.Transaction
	err := r.db.Preload("Trip.Country").Preload("User").Find(&Transac).Error

	return Transac, err
}

func (r *TransactionRepository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Trip.Country").Preload("User").Find(&transaction, ID).Error

	return transaction, err
}

func (r *TransactionRepository) DeleteTransaction(transac models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transac).Error

	return transac, err
}

// func (r *TransactionRepository) UpdateTransaction(status string, ID int) (models.Transaction, error) {
// 	var transaction models.Transaction
// 	r.db.Preload("Product").First(&transaction, ID)

// 	// If is different & Status is "success" decrement product quantity
// 	if status != transaction.Status && status == "success" {
// 		var trip models.Trip
// 		r.db.First(&trip, transaction.Trip.ID)
// 		trip.Quota = trip.Quota - 1
// 		r.db.Save(&trip)
// 	}

// 	transaction.Status = status

// 	err := r.db.Save(&transaction).Error

// 	return transaction, err
// }

func (r *TransactionRepository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Product").Preload("Product.User").Preload("Buyer").Preload("Seller").First(&transaction, "id = ?", ID).Error

	return transaction, err
}

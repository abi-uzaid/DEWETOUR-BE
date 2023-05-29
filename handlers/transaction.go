package handlers

import (
	resultdto "dewetour/dto/result"
	transactiondto "dewetour/dto/transaction"
	"dewetour/models"
	"dewetour/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type HandleTransac struct {
	TransactionRepo repositories.TransactionRepo
}

func NewHandleTransac(transactionRepo repositories.TransactionRepo) *HandleTransac {
	return &HandleTransac{TransactionRepo: transactionRepo}
}

func (h *HandleTransac) GetTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepo.GetTransaction(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	response := resultdto.SuccessResult{Code: http.StatusOK, Data: transaction}
	c.JSON(http.StatusOK, response)
}

func (h *HandleTransac) CreateTransaction(c *gin.Context) {
	// var request transactiondto.TransRequest
	// if err := c.ShouldBindJSON(&request); err != nil {
	// 	c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	// 	return

	dataFile := c.GetString("dataFile")

	counterqty, _ := strconv.Atoi(c.PostForm("counterqty"))
	total, _ := strconv.Atoi(c.PostForm("total"))
	tripid, _ := strconv.Atoi(c.PostForm("tripid"))
	userid, _ := strconv.Atoi(c.PostForm("userid"))

	request := transactiondto.TransRequest{
		CounterQty: counterqty,
		Total:      total,
		Status:     c.PostForm("status"),
		Attachment: dataFile,
		TripId:     tripid,
		UserId:     userid,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction := models.Transaction{
		CounterQty: request.CounterQty,
		Total:      request.Total,
		Status:     request.Status,
		Attachment: request.Attachment,
		TripId:     request.TripId,
		UserId:     request.UserId,
	}

	data, err := h.TransactionRepo.CreateTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := resultdto.SuccessResult{Code: http.StatusOK, Data: data}
	c.JSON(http.StatusOK, response)

}

// validation := validator.New()
// if err := validation.Struct(request); err != nil {
// 	c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	return
// }

func (h *HandleTransac) FindTransaction(c *gin.Context) {
	transaction, err := h.TransactionRepo.FindTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	for i, p := range transaction {
		transaction[i].Attachment = p.Attachment
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: transaction})
}

func (h *HandleTransac) DeleteTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepo.GetTransaction(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	data, err := h.TransactionRepo.DeleteTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: data})
}

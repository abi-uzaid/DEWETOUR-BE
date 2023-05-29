package transactiondto

type TransRequest struct {
	CounterQty int    `json:"counterqty" form:"counterqty" binding:"required"`
	Total      int    `json:"total" form:"total" binding:"required"`
	Status     string `json:"status" form:"status" binding:"required"`
	Attachment string `json:"attachment" form:"attachment" binding:"required"`
	TripId     int    `json:"tripid" form:"tripid" `
	UserId     int    `json:"userid" form:"userid"`
}

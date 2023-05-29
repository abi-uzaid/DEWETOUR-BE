package transactiondto

type TransResponse struct {
	ID         int    `json:"id"`
	CounterQty int    `json:"counterqty" form:"counterqty"`
	Total      int    `json:"total" form:"total"`
	Status     string `json:"status" form:"status"`
	Attachment string `json:"attachment" form:"attachment"`
	TripId     int    `json:"tripid" form:"tripid"`
	// UserId     int    `json:"user_id" form:"user_id"`
}

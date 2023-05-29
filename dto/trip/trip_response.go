package tripdto

type TripResponse struct {
	ID           int    `json:"id"`
	CountryId    int    `json:"country_id" form:"country_id" binding:"required"`
	Title        string `json:"title" form:"title" binding:"required"`
	Accomodation string `json:"accomodation" form:"accomodation" binding:"required"`
	Transport    string `json:"transport" form:"transport" binding:"required"`
	Eat          string `json:"eat" form:"eat" binding:"required"`
	Day          int    `json:"day" form:"day" binding:"required"`
	Night        int    `json:"night" form:"night" binding:"required"`
	Date         string `json:"date" form:"date" binding:"required"`
	Price        int    `json:"price" form:"price" binding:"required"`
	Quota        int    `json:"quota" form:"quota" binding:"required"`
	Description  string `json:"description" form:"description" binding:"required"`
	Image        string `json:"image" form:"image" binding:"required"`
}

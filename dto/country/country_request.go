package countrydto

type CountryRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"type: varchar(250)"`
}

type UpdateCountryReq struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"type: varchar(250)"`
}

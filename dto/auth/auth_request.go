package authdto

type LoginRequest struct {
	Email    string `json:"email" binding:"required" gorm:"type:varchar(255)"`
	Password string `json:"password" binding:"required" gorm:"type:varchar(255)"`
}

type RegisterRequest struct {
	Fullname string `json:"fullname" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Phone    string `json:"phone" gorm:"type:varchar(255)"`
	Address  string `json:"address" gorm:"type:varchar(255)"`
	Role     string `json:"role" gorm:"type:varchar(255)"`
}

package domain

type User struct {
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Email     *string `json:"email"`
	Phone     string  `json:"phone"  gorm:"unique;not null"`
	Password  string  `json:"password"`
	IsAdmin   bool    `json:"is_admin"`
	Model
}

package requests

type SignUpRequest struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" `
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required"`
	IsAdmin   bool   `json:"is_admin"`
}

type UpdateUserRequest struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
}

type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"Password"`
}

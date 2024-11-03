package domain

type Order struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Status    string `json:"status"`
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

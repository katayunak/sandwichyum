package model

type Sandwich struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}
type Order struct {
	ID           int    `json:"id"`
	SandwichName string `json:"sandwich_name" binding:"required"`
	TotalPrice   int    `json:"total_price"`
	User         User   `json:"user" binding:"required"`
}
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Type      string `json:"type"`
}

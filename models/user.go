package models

type User struct {
	ID           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Password     string `json:"password"`
	PasswordHash string `json:"password_hash"`
	Email        string `json:"email" binding:"required,email"`
}

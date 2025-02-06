package dto

type UserDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Platform  string `json:"platform"`
}

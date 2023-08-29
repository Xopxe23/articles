package domain

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required,gte=2"`
	Surname  string `json:"surname" binding:"required,gte=2"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

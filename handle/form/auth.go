package form

type RegisterForm struct {
	Username string `json:"username" validate:"required,min=5,max=64"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
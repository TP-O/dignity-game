package dto

type LoginPlayer struct {
	Email    string `json:"email" validate:"required,email,max=320"`
	Username string `json:"username" validate:"required,max=50"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterPlayer struct {
	Email                string `json:"email" validate:"required,email,max=320"`
	Password             string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"eqfield=Password"`
}

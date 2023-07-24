package dto

type SendResetPasswordEmail struct {
	Email string `json:"email" validate:"required,email,max=320"`
}

type ResetPassword struct {
	Password             string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"eqfield=Password"`
}

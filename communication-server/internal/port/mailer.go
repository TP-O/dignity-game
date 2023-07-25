package port

type Mailer interface {
	SendEmailVerificationEmail(to, link string) error
	SendResetPasswordEmail(to, link string) error
}

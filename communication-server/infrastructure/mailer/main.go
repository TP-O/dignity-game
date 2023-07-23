package mailer

import (
	"communication-server/config"
	"communication-server/internal/port"
	"log"

	"gopkg.in/gomail.v2"
)

type mailer struct {
	sender string
	dialer *gomail.Dialer
	closer gomail.SendCloser
}

var _ port.Mailer = (*mailer)(nil)

func New(cfg config.Mail) port.Mailer {
	var (
		dialer *gomail.Dialer
		closer gomail.SendCloser
		err    error
	)

	dialer = gomail.NewDialer(cfg.Host, cfg.Port, cfg.Address, cfg.Password)
	if closer, err = dialer.Dial(); err != nil {
		log.Panic(err)
	}

	return &mailer{
		sender: cfg.Address,
		dialer: dialer,
		closer: closer,
	}
}

func (m *mailer) Close() {
	m.closer.Close()
}

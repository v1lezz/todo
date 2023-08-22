package email

import "net/smtp"

type EmailSmtp struct {
	auth smtp.Auth
	cfg  Config
}

type Config struct {
	email    string
	password string
	host     string
	port     string
}

func NewConfig(email, password, host, port string) Config {
	return Config{email, password, host, port}
}

func NewEmailSmtp(cfg Config) *EmailSmtp {
	return &EmailSmtp{
		auth: smtp.PlainAuth("", cfg.email, cfg.password, cfg.host),
		cfg:  cfg,
	}
}

func (email *EmailSmtp) SendMessage(message, emailTo string) error {
	return smtp.SendMail(email.cfg.host+":"+email.cfg.port, email.auth, email.cfg.email, []string{emailTo}, []byte(message))
}

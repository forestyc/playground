package message

import (
	"errors"

	"gopkg.in/gomail.v2"
)

type Mail struct {
	Host     string
	Port     int
	User     string
	Password string
	Nickname string
}

func NewMail(host, user, password string, port int) *Mail {
	return &Mail{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
	}
}

// Send sends mail
func (m *Mail) Send(mailAddress []string, title, body, attachment string) error {
	if len(mailAddress) <= 0 {
		return errors.New("unspecified email address")
	}

	message := gomail.NewMessage()
	message.SetHeader("From", m.Nickname+"<"+m.User+">")
	message.SetHeader("To", mailAddress...)
	message.SetHeader("Subject", title)

	message.SetBody("text/html", body)

	if attachment != "" {
		message.Attach(attachment)
	}

	return gomail.NewDialer(m.Host, m.Port, m.User, m.Password).DialAndSend(message)
}

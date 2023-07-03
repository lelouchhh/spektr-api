package mail

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

var (
	SMTP_HOST = "smtp.mail.ru"
	SMTP_PORT = 587
)

type Sender struct {
	Sender string
	Pass   string
}

func InitSender(sender, senderPass string) Sender {

	return Sender{
		Sender: sender,
		Pass:   senderPass,
	}
}

func (s *Sender) Send(receiver, msg string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", s.Sender)
	fmt.Println("to: ", receiver)
	m.SetHeader("To", receiver)

	m.SetHeader("Subject", "Подтверждение почты!")

	m.SetBody("text/plain", msg)

	d := gomail.NewDialer(SMTP_HOST, SMTP_PORT, s.Sender, s.Pass)

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return nil
}

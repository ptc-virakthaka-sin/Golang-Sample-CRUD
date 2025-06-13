package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"learn-fiber/config"
	"log"
	"net/smtp"
)

type Request struct {
	Username string
	Subject  string
	Link     string
	Body     string
	to       string
}

func NewRequest(to, subject, body, link, username string) *Request {
	return &Request{
		Username: username,
		Subject:  subject,
		Link:     link,
		Body:     body,
		to:       to,
	}
}

func (h *Request) parseTemplate(fileName string) error {
	temp, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	var buffer bytes.Buffer
	if err = temp.Execute(&buffer, h); err != nil {
		return err
	}
	h.Body = buffer.String()
	return nil
}

func (h *Request) sendMail() error {
	addr := fmt.Sprintf("%s:%s", config.Cfg.Email.Host, config.Cfg.Email.Port)
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n"
	msg += "From: " + config.Cfg.Email.Username + "\n"
	msg += "Subject: " + h.Subject + "\n"
	msg += "To: " + h.to + "\n\n"
	msg += h.Body

	auth := smtp.PlainAuth("", config.Cfg.Email.Username, config.Cfg.Email.Password, config.Cfg.Email.Host)
	err := smtp.SendMail(addr, auth, config.Cfg.Email.Username, []string{h.to}, []byte(msg))
	return err
}

func (h *Request) Send(template string) {
	if err := h.parseTemplate(template); err != nil {
		log.Fatal(err)
	}
	if err := h.sendMail(); err != nil {
		log.Printf("Failed to send to %s\n", h.to)
		log.Print(err)
		return
	}
	log.Printf("Email sent to %s\n", h.to)
}

package consumer

import (
	"fmt"
	"learn-fiber/internal/repository"
	"learn-fiber/pkg/mail"

	"gorm.io/gorm"
)

func SendEmail(db *gorm.DB, data EmailData) {
	fmt.Println("Processing sending email", data)
	if username, err := repository.NewUser(db).GetUsername(data.To); err == nil {
		r := mail.NewRequest(data.To, data.Title, data.Body, "", username)
		r.Send("templates/change-pass.html")
	}
}

package consumer

import (
	"fmt"
	"gorm.io/gorm"
)

func SendEmail(db *gorm.DB, data interface{}) {
	fmt.Println("Processing sending email", data)
}

package consumer

import (
	"fmt"
	"gorm.io/gorm"
)

func Listener(db *gorm.DB, data map[string]interface{}) {
	if cmd, ok := data["cmd"]; ok {
		switch cmd {
		case "send_email":
			SendEmail(db, data)
		default:
			fmt.Printf("command `%s` not found", cmd)
		}
	}
}

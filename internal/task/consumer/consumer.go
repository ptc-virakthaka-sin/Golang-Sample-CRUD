package consumer

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type EmailData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	To    string `json:"to"`
}

type EmailCommand struct {
	Cmd  string    `json:"cmd"`
	Data EmailData `json:"data"`
}

func Listener(db *gorm.DB, data map[string]interface{}) {
	if cmd, err := covert(data); err == nil {
		switch cmd.Cmd {
		case "send_email":
			SendEmail(db, cmd.Data)
		default:
			fmt.Printf("command `%s` not found", cmd.Cmd)
		}
	}
}

func covert(data map[string]interface{}) (cmd EmailCommand, err error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return cmd, err
	}
	_ = json.Unmarshal([]byte(data["data"].(string)), &cmd.Data)
	_ = json.Unmarshal(bytes, &cmd)
	return cmd, nil
}

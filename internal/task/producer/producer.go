package producer

import (
	"encoding/json"
	"learn-fiber/pkg/redis"
)

func Send(payload map[string]interface{}) error {
	if data, ok := payload["data"]; ok {
		body, err := json.Marshal(data)
		if err != nil {
			return err
		}
		payload["data"] = string(body)
	}
	return redis.EnqueueTask(payload).Err()
}

package producer

import (
	"encoding/json"
	"learn-fiber/internal/task/consumer"
	"learn-fiber/pkg/redis"
)

func Send(payload consumer.EmailCommand) {
	if result := toMap(payload); result != nil {
		if body, err := json.Marshal(payload.Data); err == nil {
			result["data"] = string(body)
			redis.EnqueueTask(result)
		}
	}
}

func toMap(v interface{}) (result map[string]interface{}) {
	b, _ := json.Marshal(v)
	_ = json.Unmarshal(b, &result)
	return result
}

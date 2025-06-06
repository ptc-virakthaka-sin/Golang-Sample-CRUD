package redis

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"learn-fiber/internal/task/consumer"
	"time"

	"github.com/redis/go-redis/v9"
)

var WorkerGroup = "worker-group"
var StreamName = "task_stream"

func EnqueueTask(payload map[string]interface{}) *redis.StringCmd {
	return Client.XAdd(ctx, &redis.XAddArgs{
		Stream: StreamName,
		Values: payload,
		MaxLen: 10000,
		Approx: true,
	})
}

func StartWorker(db *gorm.DB, consumer string) {
	_ = Client.XGroupCreateMkStream(ctx, StreamName, WorkerGroup, "0").Err()
	for {
		streams, err := Client.XReadGroup(ctx, &redis.XReadGroupArgs{
			Streams:  []string{StreamName, ">"},
			Block:    5 * time.Second,
			Group:    WorkerGroup,
			Consumer: consumer,
			Count:    10,
		}).Result()

		if err != nil && !errors.Is(err, redis.Nil) {
			fmt.Println("Read error:", err)
			continue
		}
		for _, stream := range streams {
			for _, msg := range stream.Messages {
				go process(db, msg)
			}
		}
	}
}

func process(db *gorm.DB, msg redis.XMessage) {
	Client.XAck(ctx, StreamName, WorkerGroup, msg.ID)
	consumer.Listener(db, msg.Values)
}

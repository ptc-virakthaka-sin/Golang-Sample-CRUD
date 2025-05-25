package redis

import (
	"context"
	"fmt"
	"learn-fiber/config"
	"log"
	"time"

	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
)

const RedisPrefix = "core"

var Client *redis.Client

func Init() error {
	log.Println("redis is connecting...")
	addr := config.Cfg.Redis.Addr
	if addr == "" {
		addr = ":6379"
	} else {
		addr = fmt.Sprintf("%s:%s", config.Cfg.Redis.Addr, config.Cfg.Redis.Port)
	}
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Cfg.Redis.Password,
	})
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis ping err:", err)
		return err
	}
	log.Println("redis is connected")
	return nil
}

func prependPrefixKey(key string) string {
	prefix := getPrefixKeyByEnv(config.Cfg.Env)
	return fmt.Sprintf("%s:%s", prefix, key)
}

/*
env: "" | "local" | "localhost" | "dev" | "development" => prefix is "dev"
env: "sit" => prefix is "sit"
env: "uat" => prefix is "uat"
env: "prod" | "production" => prefix is ""
*/
func getPrefixKeyByEnv(env string) string {
	switch env {
	case "", "local", "localhost", "dev", "development":
		return RedisPrefix + ":dev"
	case "sit":
		return RedisPrefix + ":sit"
	case "uat":
		return RedisPrefix + ":uat"
	}
	return RedisPrefix
}

func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error) {
	var val string

	if s, ok := value.(string); ok {
		val = s
	} else {
		// convert struct to json string
		s, err := sonic.MarshalString(value)
		if err != nil {
			return "", err
		}
		val = s
	}

	resultVal, err := Client.Set(ctx, prependPrefixKey(key), val, expiration).Result()
	if err != nil {
		return "", err
	}
	return resultVal, nil
}

func Get(ctx context.Context, key string) (string, error) {
	return Client.Get(ctx, prependPrefixKey(key)).Result()
}

func Del(ctx context.Context, key string) (int64, error) {
	return Client.Del(ctx, prependPrefixKey(key)).Result()
}

func DelMulti(ctx context.Context, keys ...string) (int64, error) {
	return Client.Del(ctx, keys...).Result()
}

func Parse(ctx context.Context, key string, dest interface{}) error {
	value, err := Client.Get(ctx, prependPrefixKey(key)).Result()
	if err != nil {
		return err
	}

	err = sonic.UnmarshalString(value, dest)
	if err != nil {
		return err
	}
	return nil
}

func HSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return Client.HSet(ctx, prependPrefixKey(key)).Result()
}

func HGetAll(ctx context.Context, key string, dest interface{}) error {
	return Client.HGetAll(ctx, "key").Scan(&dest)
}

func IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return Client.IncrBy(ctx, prependPrefixKey(key), value).Result()
}

func Keys(ctx context.Context, pattern string) ([]string, error) {
	return Client.Keys(ctx, prependPrefixKey(pattern)).Result()
}

func Exists(ctx context.Context, key string) (bool, error) {
	val, err := Client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return val > 0, nil
}

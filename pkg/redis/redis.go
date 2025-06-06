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

var Prefix = "sample"
var Client *redis.Client
var ctx = context.Background()

func New() error {
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
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		log.Println("redis ping err:", err)
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
		return Prefix + ":dev"
	case "sit":
		return Prefix + ":sit"
	case "uat":
		return Prefix + ":uat"
	}
	return Prefix
}

func Set(key string, value interface{}, exp time.Duration) (string, error) {
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

	resultVal, err := Client.Set(ctx, prependPrefixKey(key), val, exp).Result()
	if err != nil {
		return "", err
	}
	return resultVal, nil
}

func Get(key string) (string, error) {
	return Client.Get(ctx, prependPrefixKey(key)).Result()
}

func Del(key string) (int64, error) {
	return Client.Del(ctx, prependPrefixKey(key)).Result()
}

func DelMulti(keys ...string) (int64, error) {
	return Client.Del(ctx, keys...).Result()
}

func Parse(key string, dest interface{}) error {
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

func HSet(key string, values ...interface{}) (int64, error) {
	return Client.HSet(ctx, prependPrefixKey(key)).Result()
}

func HGetAll(key string, dest interface{}) error {
	return Client.HGetAll(ctx, "key").Scan(&dest)
}

func IncrBy(key string, value int64) (int64, error) {
	return Client.IncrBy(ctx, prependPrefixKey(key), value).Result()
}

func Incr(key string) (int64, error) {
	return Client.Incr(ctx, prependPrefixKey(key)).Result()
}

func Exp(key string, exp time.Duration) *redis.BoolCmd {
	return Client.Expire(ctx, prependPrefixKey(key), exp)
}

func Keys(pattern string) ([]string, error) {
	return Client.Keys(ctx, prependPrefixKey(pattern)).Result()
}

func Exists(key string) (bool, error) {
	val, err := Client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return val > 0, nil
}

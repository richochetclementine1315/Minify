package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

// Context for Redis operations
// Context is used to manage deadlines, cancelation signals, and other
// request-scoped values across API boundaries and between processes.
var Ctx = context.Background()

func CreateClient(dbNum int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDRESS"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       dbNum,
	})
	return rdb
}

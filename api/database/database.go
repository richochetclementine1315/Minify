package database

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

// Context for Redis operations
// Context is used to manage deadlines, cancelation signals, and other
// request-scoped values across API boundaries and between processes.
var Ctx = context.Background()

func CreateClient(dbNum int) *redis.Client {
	addr := os.Getenv("DB_ADDRESS")
	password := os.Getenv("DB_PASSWORD")

	// Log connection attempt (helps debug Render deployment issues)
	fmt.Printf("Connecting to Redis at: %s (DB: %d)\n", addr, dbNum)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbNum,
	})

	// Test the connection
	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		fmt.Printf("Redis connection error: %v\n", err)
	}

	return rdb
}

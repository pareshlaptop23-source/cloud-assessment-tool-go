package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var RedisCtx = context.Background()

// ConnectRedis establishes a connection to Redis using environment
// variables: REDIS_HOST, REDIS_PORT, REDIS_PASSWORD, REDIS_DB.
// Returns an error on failure so callers can handle startup accordingly.
func ConnectRedis() error {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")
	dbStr := os.Getenv("REDIS_DB")

	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "6379"
	}

	addr := host + ":" + port

	db := 0
	if dbStr != "" {
		if parsed, err := strconv.Atoi(dbStr); err == nil {
			db = parsed
		}
	}

	fmt.Println("Redis connecting to:", addr)

	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           db,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	// Use a short timeout for the initial ping so startup doesn't hang.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		return fmt.Errorf("could not connect to Redis at %s: %w", addr, err)
	}

	RedisClient = client
	RedisCtx = context.Background()

	fmt.Println("Redis connected successfully:", addr)
	return nil
}

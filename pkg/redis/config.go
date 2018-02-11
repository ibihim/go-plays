package redis

import (
	"github.com/go-redis/redis"
	"os"
)

// Setup initializes a redis client
func Setup() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // use default DB
	})

	_, err := client.Ping().Result()

	return client, err
}

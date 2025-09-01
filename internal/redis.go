package internal

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewRedisClient() *redis.Client {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatalf("errore parsing redis url: %v", err)
	}

	client := redis.NewClient(opt)

	// Test connection
	if err := client.Ping(Ctx).Err(); err != nil {
		log.Fatalf("errore connessione redis: %v", err)
	}

	return client
}

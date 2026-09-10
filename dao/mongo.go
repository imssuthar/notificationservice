package dao

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client

func envOr(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

// Connect opens the shared client and verifies the server is reachable.
func Connect(ctx context.Context) error {
	c, err := mongo.Connect(options.Client().ApplyURI(envOr("MONGO_URI", "mongodb://localhost:27017")))
	if err != nil {
		return fmt.Errorf("connect mongo: %w", err)
	}
	pingCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := c.Ping(pingCtx, nil); err != nil {
		return fmt.Errorf("ping mongo: %w", err)
	}
	client = c
	return nil
}

func Disconnect(ctx context.Context) error {
	if client == nil {
		return nil
	}
	return client.Disconnect(ctx)
}

func notificationCollection() *mongo.Collection {
	return client.Database(envOr("MONGO_DB", "notificationsystem")).Collection("notifications")
}

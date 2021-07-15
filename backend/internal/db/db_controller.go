package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type DbController struct {
	Rdb *redis.Client
}

func NewDbController(endpoint string) *DbController {
	rdb := redis.NewClient(&redis.Options{
		Addr:     endpoint,
		Password: "",
		DB:       0,
	})
	return &DbController{
		Rdb: rdb,
	}
}

func (db *DbController) LoadVersion(ctx context.Context) {
	db.Rdb.Set(ctx, "version", "v0.1.0", 0)
}

func (db *DbController) LoadDefaultData(ctx context.Context) {

}

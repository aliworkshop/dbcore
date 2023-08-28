package dbcore

import (
	"context"
	"github.com/aliworkshop/error"
	"time"
)

type Cache interface {
	GetDB() any

	Store(ctx context.Context, key string, value any,
		expiration ...time.Duration) error.ErrorModel
	Expire(ctx context.Context, key string, expiration time.Duration) error.ErrorModel
	Lock(ctx context.Context, key string, expiration time.Duration) error.ErrorModel
	ListKeys(ctx context.Context, pattern string) ([]string, error.ErrorModel)
	Fetch(ctx context.Context, key string) ([]byte, error.ErrorModel)
	Load(ctx context.Context, key string, result any) error.ErrorModel
	Count(ctx context.Context, pattern string) (int, error.ErrorModel)
	CountItems(ctx context.Context, sourceKey string) (int, error.ErrorModel)
	GetExpiration(ctx context.Context, key string) (time.Duration, error.ErrorModel)
	Delete(ctx context.Context, key string) error.ErrorModel
	DeleteItem(ctx context.Context, sourceKey string, id int64) error.ErrorModel

	SetItem(ctx context.Context, key, id string, value any) error.ErrorModel
}

package dbcore

import (
	"context"
	"github.com/aliworkshop/error"
	"time"
)

type Cache interface {
	Initialize() error.ErrorModel
	GetDB() any
	Ping(ctx context.Context) error.ErrorModel

	Store(ctx context.Context, key string, value any,
		expiration ...time.Duration) error.ErrorModel
	Expire(ctx context.Context, key string, expiration time.Duration) error.ErrorModel
	Lock(ctx context.Context, key string, expiration time.Duration) error.ErrorModel
	ListKeys(ctx context.Context, pattern string) ([]string, error.ErrorModel)
	Fetch(ctx context.Context, key string) ([]byte, error.ErrorModel)
	Load(ctx context.Context, key string, result any) error.ErrorModel
	Count(ctx context.Context, pattern string) (uint64, error.ErrorModel)
	Exists(ctx context.Context, key string) (bool, error.ErrorModel)
	GetExpiration(ctx context.Context, key string) (time.Duration, error.ErrorModel)
	Delete(ctx context.Context, key string) error.ErrorModel
}

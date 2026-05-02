package dbcore

import (
	"context"
	errors "github.com/aliworkshop/error"
	"time"
)

type Cache interface {
	PubSub
	Initialize() errors.ErrorModel
	GetDB() any
	Ping(ctx context.Context) errors.ErrorModel

	Store(ctx context.Context, key string, value any, expiration ...time.Duration) errors.ErrorModel
	Expire(ctx context.Context, key string, expiration time.Duration) errors.ErrorModel
	Lock(ctx context.Context, key string, expiration time.Duration) (bool, errors.ErrorModel)
	Unlock(ctx context.Context, key string) errors.ErrorModel
	ListKeys(ctx context.Context, pattern string) ([]string, errors.ErrorModel)
	Fetch(ctx context.Context, key string) ([]byte, errors.ErrorModel)
	Load(ctx context.Context, key string, result any) errors.ErrorModel
	Count(ctx context.Context, pattern string) (uint64, errors.ErrorModel)
	Exists(ctx context.Context, key string) (bool, errors.ErrorModel)
	GetExpiration(ctx context.Context, key string) (time.Duration, errors.ErrorModel)
	Delete(ctx context.Context, key string) errors.ErrorModel
	DeleteWithPattern(ctx context.Context, pattern string) errors.ErrorModel
	Watch(ctx context.Context, key string, fn func(cache Cache) errors.ErrorModel) errors.ErrorModel
	IncrBy(ctx context.Context, key string, decrement int64) errors.ErrorModel
	DecrBy(ctx context.Context, key string, decrement int64) errors.ErrorModel
	GetInt(ctx context.Context, key string) (int64, errors.ErrorModel)
	Exec(ctx context.Context) errors.ErrorModel
	ZAdd(ctx context.Context, key string, score float64, member any) errors.ErrorModel
	ZRangeLoad(ctx context.Context, key string, min, max string, result any, offset, count int64) (any, errors.ErrorModel)
	ZRemove(ctx context.Context, key string, min, max string) errors.ErrorModel
	ZRemoveByRank(ctx context.Context, key string, from, to int64) errors.ErrorModel
	ZExist(ctx context.Context, key, member string) (bool, errors.ErrorModel)
	ZMaxScore(ctx context.Context, key string) (float64, errors.ErrorModel)
}

type PubSub interface {
	Publish(ctx context.Context, channel string, message any) errors.ErrorModel
	Subscribe(ctx context.Context, channels ...string) <-chan *redis.Message
}

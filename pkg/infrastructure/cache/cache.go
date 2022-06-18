//go:generate mockgen -destination=../../../test/mock/infrastructure/cache.go -package=infrastructure -source=$GOFILE infrastructure

package cache

import (
	"context"
	"time"
)

var (
	OneMinute = time.Minute * 1
	OneHour   = time.Minute * 60
	OneDay    = time.Hour * 24
	OneWeek   = time.Hour * 24 * 7
)

type CacheInterface interface {
	Get(ctx context.Context, key interface{}, dest interface{}) (interface{}, error)
	Set(ctx context.Context, key interface{}, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, keys ...string) error
	Flush(ctx context.Context) error
	GetSet(ctx context.Context, key string, target interface{}, function func() (interface{}, error), expiration time.Duration) error
}

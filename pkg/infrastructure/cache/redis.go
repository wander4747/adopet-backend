package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v9"
)

type Redis struct {
	client *redis.Client
}

func NewRedis() *Redis {
	con := fmt.Sprintf("%s:%s",
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
	)

	r := redis.NewClient(&redis.Options{
		Addr:     con,
		Password: "",
		DB:       0,
	})
	_, err := r.Ping(context.Background()).Result()

	if err != nil {
		return nil
	}

	return &Redis{
		client: r,
	}
}

func (r Redis) Get(ctx context.Context, key interface{}, dest interface{}) (interface{}, error) {
	result, err := r.client.Get(ctx, key.(string)).Result()
	if err != nil {
		return nil, err
	}

	switch v := interface{}(result).(type) {
	case []byte:
		err = json.Unmarshal(v, dest)
	case string:
		err = json.Unmarshal([]byte(v), dest)
	}

	return dest, err
}

func (r Redis) Set(ctx context.Context, key interface{}, value interface{}, expiration time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key.(string), bytes, expiration).Err()
}

func (r Redis) Delete(ctx context.Context, keys ...string) error {
	_, err := r.client.Del(ctx, keys...).Result()
	return err
}

func (r Redis) Flush(ctx context.Context) error {
	return r.client.FlushAll(ctx).Err()
}

func (r Redis) GetSet(ctx context.Context, key string, target interface{}, function func() (interface{}, error), expiration time.Duration) error {
	res, err := r.Get(ctx, key, &target)
	if res == nil || err != nil {
		value, err := function()
		if err != nil {
			return err
		}

		err = r.Set(ctx, key, value, expiration)
		if err != nil {
			return err
		}

		return nil
	}

	return err
}

package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/patrickmn/go-cache"
)

type Local struct {
	cache *cache.Cache
}

func NewLocal() *Local {
	return &Local{cache: cache.New(1*time.Minute, 2*time.Minute)}
}

func (r *Local) Get(_ context.Context, key interface{}, _ interface{}) (interface{}, error) {
	foo, _ := r.cache.Get(key.(string))

	return foo, nil
}

func (r *Local) Set(_ context.Context, key interface{}, value interface{}, expiration time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	r.cache.Set(key.(string), bytes, expiration)

	return nil
}

func (r *Local) Delete(_ context.Context, keys ...string) error {
	r.cache.Delete(keys[0])
	return nil
}

func (r *Local) Flush(_ context.Context) error {
	r.cache.Flush()
	return nil
}

func (r *Local) GetSet(ctx context.Context, key string, target interface{},
	function func() (interface{}, error), expiration time.Duration) error {

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
	}

	return err
}

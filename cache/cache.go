package cache

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/eko/gocache/lib/v4/cache"
	bigcachestore "github.com/eko/gocache/store/bigcache/v4"
	"time"
)

func New[T any](ctx context.Context, ttl time.Duration) *cache.Cache[T] {
	bg, err := bigcache.New(ctx, bigcache.DefaultConfig(ttl))
	if err != nil {
		panic(err)
	}
	store := bigcachestore.NewBigcache(bg)
	mgr := cache.New[T](store)
	return mgr
}

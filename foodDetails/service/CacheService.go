package service

import (
	"context"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"time"
)

type CacheService struct {
	cache *cache.Cache
}

var cs *CacheService = nil

func NewCacheService(client redis.Client) CacheService {
	if cs == nil {
		cs = new(CacheService)
		cs.cache = cache.New(&cache.Options{
			Redis:      client,
			LocalCache: cache.NewTinyLFU(1000, time.Minute),
		})
	}
	return *cs
}

func (s CacheService) SetObject(key string, object interface{}) error {
	ctx := context.TODO()
	err := s.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: object,
		TTL:   time.Hour,
	})
	return err
}

func (s CacheService) GetObject(key string, object interface{}) error {
	ctx := context.TODO()
	err := s.cache.Get(ctx, key, object)
	return err
}

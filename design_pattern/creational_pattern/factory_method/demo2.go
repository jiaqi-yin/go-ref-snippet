package main

import (
	"fmt"
	"reflect"
)

type Cache interface {
	Set(key, value string)
	Get(string) string
}

type RedisCache struct {
	data map[string]string
}

func NewRedisCache() *RedisCache {
	return &RedisCache{
		data: make(map[string]string),
	}
}

func (redis *RedisCache) Set(key, value string) {
	redis.data[key] = value
}

func (redis *RedisCache) Get(key string) string {
	return fmt.Sprintf("%v: %v=%v", reflect.ValueOf(redis).Type().String(), key, redis.data[key])
}

type MemCache struct {
	data map[string]string
}

func (mem *MemCache) Set(key, value string) {
	mem.data[key] = value
}

func (mem *MemCache) Get(key string) string {
	return fmt.Sprintf("%v: %v=%v", reflect.ValueOf(mem).Type().String(), key, mem.data[key])
}

func NewMemCache() *MemCache {
	return &MemCache{
		data: make(map[string]string),
	}
}

type CacheFactory interface {
	Create() Cache
}

type RedisCacheFactory struct{}

func (rcf RedisCacheFactory) Create() Cache {
	return NewRedisCache()
}

type MemCacheFactory struct{}

func (mcf MemCacheFactory) Create() Cache {
	return NewMemCache()
}

func main() {
	redisCacheFactory := RedisCacheFactory{}
	redisCache := redisCacheFactory.Create()
	redisCache.Set("k1", "v1")
	fmt.Println(redisCache.Get("k1"))

	memCacheFactory := MemCacheFactory{}
	memCache := memCacheFactory.Create()
	memCache.Set("k2", "v2")
	fmt.Println(memCache.Get("k2"))
}

package main

import (
	"errors"
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

type cacheType int

const (
	redis cacheType = iota
	mem
)

type CacheFactory struct{}

func (factory *CacheFactory) Create(ct cacheType) (Cache, error) {
	if ct == redis {
		return &RedisCache{
			data: map[string]string{},
		}, nil
	}

	if ct == mem {
		return &MemCache{
			data: map[string]string{},
		}, nil
	}

	return nil, errors.New("error cache type")
}

func main() {
	cacheFactory := &CacheFactory{}

	redis, error := cacheFactory.Create(redis)
	if error != nil {
		panic(error)
	}
	redis.Set("k1", "v1")
	fmt.Println(redis.Get("k1"))

	mem, error := cacheFactory.Create(mem)
	if error != nil {
		panic(error)
	}
	mem.Set("k1", "v1")
	fmt.Println(mem.Get("k1"))
}

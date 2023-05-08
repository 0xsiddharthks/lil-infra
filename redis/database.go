package redis

import (
	"errors"
	"sync"
	"time"
)

type pair[T, U any] struct {
	first  T
	second U
}

type redis struct {
	umap map[string]pair[string, int64]
}

var redisInstance *redis
var once = sync.Once{}

func Instance() DB {
	if redisInstance == nil {
		once.Do(func() {
			redisInstance = &redis{umap: make(map[string]pair[string, int64])}
		})
	}
	return redisInstance
}

type DB interface {
	set(key string, value string, ttl int64)
	get(key string) (string, error)
}

func (db redis) set(key string, value string, ttl int64) {
	db.umap[key] = pair[string, int64]{first: value, second: ttl}
}

func (db redis) get(key string) (string, error) {
	if val, ok := db.umap[key]; ok {
		ttl := val.second
		if ttl == -1 || ttl > time.Now().Unix()*1000 {
			return val.first, nil
		}
		return "", errors.New("key expired")
	}
	return "", errors.New("key not found")
}

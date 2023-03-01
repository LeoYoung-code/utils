package utils

import (
	"time"
	
	"github.com/patrickmn/go-cache"
)

var cacheDriver = cache.New(5*time.Minute, 10*time.Minute)

const (
	CACHE_EXP_1DAY = time.Hour * 24
)

func SetGoCache[T any](key string, val T, exp time.Duration) {
	cacheDriver.Set(key, val, exp)
}

func GetGoCache[T any](key string, defaultVal T) (T, bool) {
	val, ok := cacheDriver.Get(key)
	if !ok {
		return defaultVal, ok
	}
	return val.(T), ok
}

func DelGoCache(key string) {
	cacheDriver.Delete(key)
}

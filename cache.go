package utils

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"github.com/patrickmn/go-cache"
)

var cacheDriver = cache.New(1*time.Minute, 10*time.Minute)

// SetGoCache set cache
func SetGoCache[T any](key string, val T, exp time.Duration) {
	cacheDriver.Set(key, val, exp)
}

// GetGoCache get cache
func GetGoCache[T any](key string, defaultVal T) (T, bool) {
	val, ok := cacheDriver.Get(key)
	if !ok {
		return defaultVal, ok
	}
	return val.(T), ok
}

// DelGoCache delete cache
func DelGoCache(key string) {
	cacheDriver.Delete(key)
}

// DeepCopy deep copy
func DeepCopy(to, form any) {
	err := copier.CopyWithOption(to, form, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
	if err != nil {
		log.Error(err)
	}
}

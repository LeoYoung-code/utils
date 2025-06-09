package cache

import (
	"fmt"
	"time"

	json "github.com/bytedance/sonic"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/jinzhu/copier"
	"github.com/patrickmn/go-cache"
)

var (
	defaultExpire = 1 * time.Minute
	defaultPurge  = 10 * time.Minute
	cacheDriver   *cache.Cache
)

func init() {
	cacheDriver = cache.New(defaultExpire, defaultPurge)
}

// SetGoCache set cache
func SetGoCache[T any](key string, val T, exp time.Duration) {
	cacheDriver.Set(key, val, exp)
}

// GetGoCache get cache with safe type assertion
func GetGoCache[T any](key string, defaultVal T) (T, bool) {
	val, ok := cacheDriver.Get(key)
	if !ok {
		return defaultVal, false
	}

	result, ok := val.(T)
	if !ok {
		log.Warnf("cache type assertion failed for key %s, expected type %T, got %T", key, result, val)
		return defaultVal, false
	}
	return result, true
}

// SetGoCacheWithDeep 写入缓存并返回错误
func SetGoCacheWithDeep[T any](key string, val T, exp time.Duration) error {
	marshal, err := json.Marshal(val)
	if err != nil {
		log.Error(err)
		return err
	}
	cacheDriver.Set(key, marshal, exp)
	return nil
}

// GetGoCacheWithDeep 只能用指针结构取数据
func GetGoCacheWithDeep(key string, to any) error {
	val, ok := cacheDriver.Get(key)
	if !ok {
		return fmt.Errorf("cache miss for key %s", key)
	}

	data, ok := val.([]byte)
	if !ok {
		return fmt.Errorf("invalid cached data type for key %s, expected []byte, got %T", key, val)
	}

	if err := json.Unmarshal(data, to); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// DeepCopy deep copy with error return
func DeepCopy(to, form any) error {
	opt := copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	}

	if err := copier.CopyWithOption(to, form, opt); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

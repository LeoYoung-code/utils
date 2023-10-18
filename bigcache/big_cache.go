package bigcache

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/go-kratos/kratos/v2/log"
)

var config = bigcache.Config{
	// number of shards (must be a power of 2)
	Shards: 1024,

	// time after which entry can be evicted
	LifeWindow: 1 * time.Minute,

	// Interval between removing expired entries (clean up).
	// If set to <= 0 then no action is performed.
	// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
	CleanWindow: 5 * time.Minute,

	// rps * lifeWindow, used only in initial memory allocation
	MaxEntriesInWindow: 1000 * 10 * 100,

	// max entry size in bytes, used only in initial memory allocation
	MaxEntrySize: 1000,

	// prints information about additional memory allocation
	Verbose: false,

	// cache will not allocate more memory than this limit, value in MB
	// if value is reached then the oldest entries can be overridden for the new ones
	// 0 value means no size limit
	HardMaxCacheSize: 8192,

	// callback fired when the oldest entry is removed because of its expiration time or no space left
	// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
	// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
	OnRemove: nil,

	// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
	// for the new entry, or because delete was called. A constant representing the reason will be passed through.
	// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
	// Ignored if OnRemove is specified.
	OnRemoveWithReason: nil,
}

func NewCache(ctx context.Context) *bigcache.BigCache {
	cache, err := bigcache.New(ctx, config)
	if err != nil {
		log.Fatal("bigCache 初始化失败", err)
	}
	return cache
}

var cacheDriver = NewCache(context.Background())

// GetGoCache 只获取是否存在缓存的状态
func GetGoCache(key string) bool {
	_, err := cacheDriver.Get(key)
	if err != nil {
		if err != bigcache.ErrEntryNotFound {
			log.Error("Get报错 key：", key, " 错误： ", err)
		}
		return false
	}
	return true
}

func SetGoCache(key string, val []byte) {
	err := cacheDriver.Set(key, val)
	if err != nil {
		log.Error("Set报错 key：", key, " 错误： ", err)
	}
}

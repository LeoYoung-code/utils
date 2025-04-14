package bigcache

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/go-kratos/kratos/v2/log"
)

// Config 是缓存配置，允许外部自定义
type Config struct {
	// 分片数量（必须是2的幂）
	Shards int
	// 条目可被淘汰的时间
	LifeWindow time.Duration
	// 清理过期条目的时间间隔
	CleanWindow time.Duration
	// 窗口中的最大条目数，用于初始内存分配
	MaxEntriesInWindow int
	// 最大条目大小（字节），用于初始内存分配
	MaxEntrySize int
	// 是否输出详细日志
	Verbose bool
	// 缓存最大内存限制（MB），0表示无限制
	HardMaxCacheSize int
	// 条目被移除时的回调函数
	OnRemoveWithReason func(key string, entry []byte, reason bigcache.RemoveReason)
}

// DefaultConfig 返回默认配置
func DefaultConfig() Config {
	return Config{
		Shards:             1024,
		LifeWindow:         1 * time.Minute,
		CleanWindow:        5 * time.Minute,
		MaxEntriesInWindow: 1000 * 10 * 100,
		MaxEntrySize:       1000,
		Verbose:            false,
		HardMaxCacheSize:   8192,
		OnRemoveWithReason: nil,
	}
}

// Cache 是对bigcache的封装
type Cache struct {
	client *bigcache.BigCache
	logger log.Logger
}

// NewCache 创建新的缓存实例
func NewCache(ctx context.Context, cfg Config, logger log.Logger) (*Cache, error) {
	if logger == nil {
		logger = log.DefaultLogger
	}

	bigcacheCfg := bigcache.Config{
		Shards:             cfg.Shards,
		LifeWindow:         cfg.LifeWindow,
		CleanWindow:        cfg.CleanWindow,
		MaxEntriesInWindow: cfg.MaxEntriesInWindow,
		MaxEntrySize:       cfg.MaxEntrySize,
		Verbose:            cfg.Verbose,
		HardMaxCacheSize:   cfg.HardMaxCacheSize,
		OnRemoveWithReason: cfg.OnRemoveWithReason,
	}

	cache, err := bigcache.New(ctx, bigcacheCfg)
	if err != nil {
		return nil, err
	}

	return &Cache{
		client: cache,
		logger: logger,
	}, nil
}

// 全局缓存实例
var defaultCache *Cache

// InitDefaultCache 初始化默认缓存
func InitDefaultCache(ctx context.Context, cfg Config, logger log.Logger) error {
	cache, err := NewCache(ctx, cfg, logger)
	if err != nil {
		return err
	}
	defaultCache = cache
	return nil
}

// MustInitDefaultCache 初始化默认缓存，失败则触发panic
func MustInitDefaultCache(ctx context.Context, cfg Config, logger log.Logger) {
	if err := InitDefaultCache(ctx, cfg, logger); err != nil {
		panic(err)
	}
}

// GetOrInit 获取默认缓存，如果未初始化则使用默认配置初始化
func GetOrInit(ctx context.Context) *Cache {
	if defaultCache == nil {
		cache, err := NewCache(ctx, DefaultConfig(), log.DefaultLogger)
		if err != nil {
			log.Fatal(err)
		}
		defaultCache = cache
	}
	return defaultCache
}

// Get 从缓存获取值
func (c *Cache) Get(key string) ([]byte, error) {
	return c.client.Get(key)
}

// Set 设置缓存
func (c *Cache) Set(key string, val []byte) error {
	return c.client.Set(key, val)
}

// Delete 删除缓存
func (c *Cache) Delete(key string) error {
	return c.client.Delete(key)
}

// Len 返回缓存条目数量
func (c *Cache) Len() int {
	return c.client.Len()
}

// Reset 重置缓存
func (c *Cache) Reset() error {
	return c.client.Reset()
}

// Close 关闭缓存
func (c *Cache) Close() error {
	return c.client.Close()
}

// GetFromDefaultCache 从默认缓存中获取值
func GetFromDefaultCache(key string) ([]byte, bool) {
	cache := GetOrInit(context.Background())
	val, err := cache.Get(key)
	if err != nil {
		if err != bigcache.ErrEntryNotFound {
			cache.logger.Log(log.LevelError, "msg", "获取缓存失败", "key", key, "err", err)
		}
		return nil, false
	}
	return val, true
}

// SetToDefaultCache 设置值到默认缓存
func SetToDefaultCache(key string, val []byte) bool {
	cache := GetOrInit(context.Background())
	err := cache.Set(key, val)
	if err != nil {
		cache.logger.Log(log.LevelError, "msg", "设置缓存失败", "key", key, "err", err)
		return false
	}
	return true
}

// Exists 检查键是否存在于默认缓存
func Exists(key string) bool {
	cache := GetOrInit(context.Background())
	_, err := cache.Get(key)
	return err == nil
}

// DeleteFromDefaultCache 从默认缓存删除键
func DeleteFromDefaultCache(key string) bool {
	cache := GetOrInit(context.Background())
	err := cache.Delete(key)
	if err != nil && err != bigcache.ErrEntryNotFound {
		cache.logger.Log(log.LevelError, "msg", "删除缓存失败", "key", key, "err", err)
		return false
	}
	return true
}

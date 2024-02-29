package jgtime

import (
	"sync"
	"time"
)

// LocationCache 用于存储和重用 time.Location 对象
type LocationCache struct {
	mu       sync.RWMutex
	location map[string]*time.Location
}

// NewLocationCache 创建一个新的 LocationCache 实例
func NewLocationCache() *LocationCache {
	return &LocationCache{
		location: make(map[string]*time.Location),
	}
}

// LoadLocation 从缓存中获取 time.Location 对象，如果不存在则加载
func (lc *LocationCache) LoadLocation(name string) (*time.Location, error) {
	lc.mu.RLock()
	loc, ok := lc.location[name]
	lc.mu.RUnlock()

	if ok {
		// 时区已经在缓存中，直接返回
		return loc, nil
	}

	// 时区不在缓存中，加载并存入缓存
	loc, err := time.LoadLocation(name)
	if err != nil {
		return nil, err
	}

	lc.mu.Lock()
	lc.location[name] = loc
	lc.mu.Unlock()

	return loc, nil
}

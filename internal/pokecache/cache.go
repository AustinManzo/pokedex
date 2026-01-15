package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu    sync.RWMutex
	entry map[string]cacheEntry
}
type cacheEntry struct {
	createdAt time.Time
	data      []byte
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	newEntry := cacheEntry{
		createdAt: time.Now(),
		data:      value,
	}
	c.entry[key] = newEntry
	c.mu.Unlock()

}
func (c *Cache) Get(key string) ([]byte, bool) {

	c.mu.RLock()
	entry, ok := c.entry[key]
	if !ok {
		c.mu.RUnlock()
		return nil, false
	}
	val := entry.data
	c.mu.RUnlock()
	return val, true

}
func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.entry {
		age := time.Since(v.createdAt)
		if age > interval {
			delete(c.entry, k)
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entry: make(map[string]cacheEntry),
	}

	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			cache.reapLoop(interval)
		}
	}()

	return cache
}

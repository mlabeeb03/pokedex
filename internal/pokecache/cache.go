package pokecache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]CacheEntry
	mu    *sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]CacheEntry),
		mu:    &sync.RWMutex{},
	}
	go c.ReapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = CacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
	log.Printf("Cached Key %s", key)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	log.Printf("Returned cached Key %s", key)
	return val.val, true
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		for k, v := range c.cache {
			if v.createdAt.Before(time.Now().Add(-interval)) {
				c.mu.Lock()
				delete(c.cache, k)
				c.mu.Unlock()
				log.Printf("\nDeleted cached Key %s", k)
				fmt.Printf("Pokedex > ")
			}
		}
	}
}

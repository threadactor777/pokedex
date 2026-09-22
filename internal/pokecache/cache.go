package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu      sync.Mutex
	entries map[string]cacheEntry
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.entries[key]; exists {
		fmt.Println("Entry (Key) already exists!")
	} else {
		c.entries[key] = cacheEntry{
			createdAt: time.Now(),
			val:       val,
		}
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, exists := c.entries[key]; exists {
		return entry.val, true
	} else {
		return nil, false
	}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) reapLoop(interval time.Duration) {
	/*Creates a new time.Ticker that will repeatedly
	send a message (a tick) down a channel at the specified
	time frequency*/
	ticker := time.NewTicker(interval)

	/* infinite loop that blocks and waits. Every time the
	ticker triggers, a new value is sent down the ticker.C channel*/
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, entry := range c.entries {
			if now.Sub(entry.createdAt) > interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

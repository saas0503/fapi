package cache

import (
	"github.com/saas0503/futils"
	"sync"
	"time"
)

type item struct {
	exp uint32
	val any
}

type Cache struct {
	sync.RWMutex
	data map[string]item
}

func New() *Cache {
	store := &Cache{
		data: make(map[string]item),
	}

	futils.StartTimeStampUpdater()
	go store.gc(1 * time.Second)
	return store
}

// Get value by key
func (c *Cache) Get(key string) any {
	c.RLock()
	v, ok := c.data[key]
	c.RUnlock()
	if !ok || v.exp != 0 && v.exp <= futils.Timestamp() {
		return nil
	}
	return v.val
}

// Set key with value
func (c *Cache) Set(key string, val any, ttl time.Duration) {
	var exp uint32
	if ttl > 0 {
		exp = uint32(ttl.Seconds()) + futils.Timestamp()
	}
	i := item{exp, val}
	c.Lock()
	c.data[key] = i
	c.Unlock()
}

// Delete key by key
func (c *Cache) Delete(key string) {
	c.Lock()
	delete(c.data, key)
	c.Unlock()
}

// Reset all keys
func (c *Cache) Reset() {
	nd := make(map[string]item)
	c.Lock()
	c.data = nd
	c.Unlock()
}

func (c *Cache) gc(sleep time.Duration) {
	ticker := time.NewTicker(sleep)
	defer ticker.Stop()
	var expired []string

	for range ticker.C {
		ts := futils.Timestamp()
		expired = expired[:0]
		c.RLock()
		for key, v := range c.data {
			if v.exp != 0 && v.exp <= ts {
				expired = append(expired, key)
			}
		}
		c.RUnlock()
		c.Lock()
		// Double-checked locking.
		// We might have replaced the item in the meantime.
		for i := range expired {
			v := c.data[expired[i]]
			if v.exp != 0 && v.exp <= ts {
				delete(c.data, expired[i])
			}
		}
		c.Unlock()
	}
}

package cache

import (
	"sync"
	"time"

	"github.com/dgraph-io/ristretto"
)

type CacheHandler struct {
	Cache *ristretto.Cache
}

var Ch CacheHandler

func Init() {
	Ch.CreateCache()
}

func (c *CacheHandler) CreateCache() {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}

	c.Cache = cache
}

func (c *CacheHandler) Get(key string) (interface{}, bool) {
	value, found := c.Cache.Get(key)
	if !found {
		return value, found
	}

	return value, found
}

func (c *CacheHandler) Set(key string, value interface{}) bool {
	return c.Cache.Set(key, value, 0)
}

func (c *CacheHandler) Del(key string) {
	c.Cache.Del(key)
}

// Make sure key is stored in cache
func (c *CacheHandler) WaitForCacheWrite(sportId string, ch chan interface{}, mx *sync.Mutex) {
	for {
		value, ok := c.Cache.Get(sportId)

		if ok {
			ch <- value
			mx.Unlock()
			return
		}

		time.Sleep(time.Millisecond * 100)
	}
}

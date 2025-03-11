package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheValue
	lock sync.Mutex
}

type cacheValue struct {
	value      interface{}
	expiration time.Time
}

func newCache() *Cache {
	return &Cache{
		data: make(map[string]cacheValue),
	}
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()

	expirationTime := time.Now().Add(expiration)
	c.data[key] = cacheValue{
		value:      value,
		expiration: expirationTime,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	value, ok := c.data[key]
	if !ok || time.Now().After(value.expiration) {
		delete(c.data, key)
		return nil, false
	}

	return value.value, true
}

func main() {
	cache := newCache()

	cache.Set("key_1", "value1", 10*time.Second)
	cache.Set("key_2", 23, 5*time.Second)

	time.Sleep(7 * time.Second)

	value1, ok1 := cache.Get("key_1")
	value2, ok2 := cache.Get("key_2")

	fmt.Println("key_1: ", value1, ok1)
	fmt.Println("key_2 ", value2, ok2)

}

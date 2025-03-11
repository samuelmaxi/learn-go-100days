package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu    sync.RWMutex
	store map[string]cacheItem
}

type cacheItem struct {
	value      interface{}
	expiration int64
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]cacheItem),
	}
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = cacheItem{
		value:      value,
		expiration: time.Now().Add(duration).Unix(),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.store[key]
	if !found {
		return nil, false
	}

	if time.Now().Unix() > item.expiration {
		delete(c.store, key)
		return nil, false
	}

	return item.value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, key)
}

func main() {
	cache := NewCache()

	cache.Set("username", "golangUser", 5*time.Second)

	if value, found := cache.Get("username"); found {
		fmt.Println("valor encontrado no cache: ", value)
	} else {
		fmt.Println("valor nao encotrado no cache")
	}

	time.Sleep((6 * time.Second))

	if value, found := cache.Get("username"); found {
		fmt.Println("Valor encontrado no cache", value)
	} else {
		fmt.Println("Valor expirado e removido do cache")
	}

}

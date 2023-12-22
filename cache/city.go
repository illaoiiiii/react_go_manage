package cache

import "sync"

type CityCache struct {
	data map[int]interface{}
	mu   sync.RWMutex
}

func NewCityCache() *CityCache {
	return &CityCache{
		data: make(map[int]interface{}),
	}
}

func (c *CityCache) Get(key int) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.data[key]
	return value, ok
}

func (c *CityCache) FindKeyByValue(value interface{}) int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for key, val := range c.data {
		if val == value {
			return key
		}
	}

	return -1
}

func (cc *CityCache) GetAllKeyValues() map[int]interface{} {
	cc.mu.RLock()
	defer cc.mu.RUnlock()

	keyValues := make(map[int]interface{})
	for key, value := range cc.data {
		keyValues[key] = value
	}

	return keyValues
}

func (c *CityCache) Set(key int, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

func (c *CityCache) Delete(key int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}

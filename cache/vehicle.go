package cache

import "sync"

type VehicleCache struct {
	data map[int]interface{}
	mu   sync.RWMutex
}

func NewVehicleCache() *VehicleCache {
	return &VehicleCache{
		data: make(map[int]interface{}),
	}
}

func (c *VehicleCache) Get(key int) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.data[key]
	return value, ok
}

func (c *VehicleCache) FindKeyByValue(value interface{}) int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for key, val := range c.data {
		if val == value {
			return key
		}
	}

	return -1
}

func (cc *VehicleCache) GetAllKeyValues() map[int]interface{} {
	cc.mu.RLock()
	defer cc.mu.RUnlock()

	keyValues := make(map[int]interface{})
	for key, value := range cc.data {
		keyValues[key] = value
	}

	return keyValues
}

func (c *VehicleCache) Set(key int, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

func (c *VehicleCache) Delete(key int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}

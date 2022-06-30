package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.RWMutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.items[key]; ok {
		item := v.Value.(cacheItem)
		item.value = value
		v.Value = item
		c.queue.MoveToFront(v)
		return true
	}

	if c.queue.Len() == c.capacity {
		oldItem := c.queue.Back().Value.(cacheItem)
		delete(c.items, oldItem.key)
		c.queue.Remove(c.queue.Back())
	}

	newItem := cacheItem{
		key:   key,
		value: value,
	}

	c.queue.PushFront(newItem)
	c.items[key] = c.queue.Front()
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.items[key]; ok {
		c.queue.MoveToFront(v)
		item := v.Value.(cacheItem)
		return item.value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.queue.Clear()
	c.items = make(map[Key]*ListItem, c.capacity)
}

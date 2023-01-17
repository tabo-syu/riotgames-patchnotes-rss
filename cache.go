package main

import riotgames "github.com/tabo-syu/riotgames-game-articles"

type Cache[T riotgames.Article] struct {
	storage map[string][]*T
}

func NewCache[T riotgames.Article]() *Cache[T] {
	return &Cache[T]{
		storage: map[string][]*T{},
	}
}

func (c *Cache[T]) Get(key string) []*T {
	if value, ok := c.storage[key]; ok {
		return value
	}

	return nil
}

func (c *Cache[T]) Set(key string, value []*T) {
	c.storage[key] = value
}

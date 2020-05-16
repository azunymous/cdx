package gocache

import (
	"fmt"
	"github.com/azunymous/cdx/watch"
	"github.com/patrickmn/go-cache"
	"time"
)

type GoCache struct {
	c *cache.Cache
}

func NewGoCache(defaultExpiration, cleanupInterval time.Duration) *GoCache {
	c := cache.New(defaultExpiration, cleanupInterval)
	return &GoCache{c: c}
}

var _ watch.DiffStore = &GoCache{}

func (g *GoCache) Get(key string) (string, error) {
	get, ok := g.c.Get(key)
	if !ok {
		return "", fmt.Errorf("no patch with name %s found", key)
	}
	return get.(string), nil
}

func (g *GoCache) Set(key, value string) error {
	if _, ok := g.c.Get(key); ok {
		return fmt.Errorf("patch with name %s already exists", key)
	}
	g.c.Set(key, value, cache.DefaultExpiration)
	return nil
}

package cache

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCachePut ensures that Cache.Put increments the cache size.
func TestCachePut(t *testing.T) {
	key := "Foo"
	value := "Bar"
	newCache := NewCache(3)
	newCache.Put(key, value)
	assert.Equal(t, newCache.Size(), 1)
}

// TestCachePuts tests multiple cache puts and ensures that cache size stays under capacity.
func TestCachePuts(t *testing.T) {
	newCache := NewCache(3)
	newCache.Put("one", 1)
	newCache.Put("two", 2)
	newCache.Put("three", 3)
	assert.Equal(t, 3, newCache.Size())
}

// TestCacheGet ensures that Cache.Get properly retrieves an item that exists, and throws an error for an item that doesn't.
func TestCacheGet(t *testing.T) {
	key := "Foo"
	value := "Bar"
	newCache := NewCache(3)
	newCache.Put(key, value)
	item, err := newCache.Get(key)
	assert.Nil(t, err)
	assert.Equal(t, value, item)
	_, err = newCache.Get("doesn't exist")
	assert.Error(t, err)
}

// TestCacheEviction add items to the cache, and ensures that the least recently used is evicted.
func TestCacheEviction(t *testing.T) {
	newCache := NewCache(3)
	newCache.Put("one", 1)
	newCache.Put("two", 2)
	newCache.Put("three", 3)
	assert.Equal(t, 3, newCache.Size())
	assert.Equal(t, 1, newCache.LRUList.Head.Data.(Item).value)
	assert.Equal(t, 3, newCache.LRUList.Tail.Data.(Item).value)

	// Cause one to be evicted
	newCache.Put("four", 4)
	assert.Equal(t, 3, newCache.Size())
	assert.Equal(t, newCache.Size(), newCache.capacity)
	assert.Equal(t, 3, len(newCache.itemMap))
	_, ok := newCache.itemMap["one"]
	assert.False(t, ok)
	_, err := newCache.Get("one")
	assert.Error(t, err)

	// set two as most recently used
	newCache.Put("two", 2)
	assert.Equal(t, 2, newCache.LRUList.Tail.Data.(Item).value)
	fmt.Println(newCache)

	// now cause three to be evicted
	newCache.Put("five", 5)
	assert.Equal(t, 5, newCache.LRUList.Tail.Data.(Item).value)
	_, err = newCache.Get("three")
	assert.Error(t, err)
}

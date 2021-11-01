package cache

import (
	"fmt"
	"log"

	"github.com/Ben-harder/gocache/doubleLinkedList"
)

// linked list which represents the cache access order
// most recently used item is always at the head of the linked list
// Put:
// 	check map to see if the key exists.
// 	If it does, then we need to find it in the linked list.
// 		once we find it, we extract it from that spot and attach it to the end of the list
// 	If it doesn't, we can add it directly to the end of the list
//		Check capacity and snip and destroy the first item if needed
//	The value of the map could be the pointer to the node in the linked list? That would give us O(1) access and puts

// cache:
// 	map[key]*Node
//	tail *Node
// 	head *Node
//	size int
//	capacity int

// Get(CacheItem) error
// 	check map to see if key exists
// 		If it does, then pull the node pointer from the map, then the value from the node pointer. Attach the node to the end of the list.
//	otherwise, return some fail code

type Key interface{}

type CacheValue interface{}

type CacheItem struct {
	key   Key
	value CacheValue
}

type Cache struct {
	itemMap  map[Key]*doubleLinkedList.Node
	LRUList  *doubleLinkedList.List
	capacity int
}

func NewCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		itemMap:  make(map[Key]*doubleLinkedList.Node),
		LRUList:  new(doubleLinkedList.List),
	}
}

// Get will return the CacheValue given a key and an error if the item wasn't in the cache.
func (c *Cache) Get(key Key) (CacheValue, error) {
	n, ok := c.itemMap[key]
	if !ok {
		return nil, fmt.Errorf("item not found")
	}
	// set CacheItem as most recent
	c.setMostRecent(n)
	val := n.Data.(CacheItem).value
	return val, nil
}

func (c *Cache) setMostRecent(n *doubleLinkedList.Node) {
	// First remove the node, then send it to the back
	c.LRUList.Remove(n)
	c.LRUList.InsertEnd(n)
}

func (c *Cache) evict() {
	head := c.LRUList.Head
	log.Printf("evicting LRU from cache %v\n", head.Data.(CacheItem))
	c.LRUList.Remove(head)
	delete(c.itemMap, head.Data.(CacheItem).key)
}

// insert puts a new item into the cache
func (c *Cache) insert(key Key, item CacheValue) {
	newNode := &doubleLinkedList.Node{Data: CacheItem{key: key, value: item}}
	c.LRUList.InsertEnd(newNode)
	c.itemMap[key] = newNode
}

func (c *Cache) Put(key Key, item CacheValue) {
	n, ok := c.itemMap[key]
	if ok {
		// We already have the item in cache
		c.setMostRecent(n)
		return
	}

	// Check if the cache is full and evict the LRU item if it is
	if c.Size() == c.capacity {
		c.evict()
	}
	c.insert(key, item)
}

func (c *Cache) Size() int {
	return c.LRUList.Size
}

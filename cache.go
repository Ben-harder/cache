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
	tail     *doubleLinkedList.Node
	head     *doubleLinkedList.Node
	size     int
	capacity int
}

func NewCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		itemMap:  make(map[Key]*doubleLinkedList.Node),
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
	// If the cache is empty, the tail is the head is now the most recent
	if c.tail == nil {
		c.tail = n
		return
	}
	c.tail.Next = n
	n.Prev = c.tail
	c.tail = n

}

func (c *Cache) evict() {
	temp := c.head
	c.head = c.head.Next
	c.head.Prev = nil
	log.Printf("evicting LRU from cache %v\n", temp.Data.(CacheItem))
	delete(c.itemMap, temp.Data.(CacheItem).key)
	c.size -= 1
}

// insert puts a new item into the cache
func (c *Cache) insert(key Key, item CacheValue) {
	newNode := &doubleLinkedList.Node{Prev: c.tail, Data: CacheItem{key: key, value: item}}
	// If the cache is empty, the tail is the head is now the most recent
	if c.head == nil {
		c.head = newNode
	}
	c.setMostRecent(newNode)
	c.itemMap[key] = newNode
	c.size += 1
}

func (c *Cache) Put(key Key, item CacheValue) {
	n, ok := c.itemMap[key]
	if ok {
		// We already have the item in cache
		c.setMostRecent(n)
		return
	}

	// Check if the cache is full and evict the LRU item if it is
	if c.size == c.capacity {
		c.evict()
	}
	c.insert(key, item)
}

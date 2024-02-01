package lrucache

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

// Main impl struct
type lruCacheImpl[Key comparable, Value any] struct {
	// Store max size of cache
	capacity int

	// Items map to search by key
	items map[Key]valueEntity[Value]

	// Priority double linked list
	// Front of list is the highest priority (recent assess)
	// Back of list is the lowest priority (remove ASAP)
	priorities list.List

	// Mutex for thread safe
	mutex sync.Mutex
}

// Helper structs
type valueEntity[Value any] struct {
	value          Value
	priorityEntity *list.Element
}

// Constructor

func NewLRUCache[Key comparable, Value any](capacity int) LRUCache[Key, Value] {
	if capacity <= 0 {
		panic(errors.New(fmt.Sprintf("Capacity too small %d", capacity)))
	}

	impl := &lruCacheImpl[Key, Value]{
		capacity: capacity,
		items:    make(map[Key]valueEntity[Value], capacity),
	}

	impl.priorities.Init()

	return impl
}

// Implementation of interface LRUCache

var _ LRUCache[interface{}, interface{}] = (*lruCacheImpl[interface{}, interface{}])(nil)

func (lc *lruCacheImpl[Key, Value]) Add(key Key, value Value) bool {
	// Sync
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	// Check if value already exist
	if _, exist := lc.items[key]; exist {
		return false
	}

	// Entity is not exits, we can create it
	// Create new entity
	var valueEntity valueEntity[Value]
	// Set value
	valueEntity.value = value
	// Set priority entity and store it into priority list
	valueEntity.priorityEntity = lc.priorities.PushFront(key)

	// Store into map
	lc.items[key] = valueEntity

	// Check size
	if lc.len() > lc.cap() {
		// Remove one
		keyToRemove, casted := lc.priorities.Remove(lc.priorities.Back()).(Key)
		if !casted {
			panic(errors.New("Can`t cast priority item to the key"))
		}
		_, keyExist := lc.items[keyToRemove]
		if !keyExist {
			panic(errors.New("Key is not exist in the map of values"))
		}
		delete(lc.items, keyToRemove)
	}

	return true
}

func (lc *lruCacheImpl[Key, Value]) Get(key Key) (Value, bool) {
	// Sync
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	// Search item in map
	valueEntity, ok := lc.items[key]
	if !ok {
		// Return default object and false
		return valueEntity.value, ok
	}

	// Move priority entity to the front of the list
	lc.priorities.MoveToFront(valueEntity.priorityEntity)

	// Successfully found
	return valueEntity.value, true
}

func (lc *lruCacheImpl[Key, Value]) Remove(key Key) bool {
	// Sync
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	// Search and delete
	if item, ok := lc.items[key]; ok {

		// Remove from priority list
		removedKey := lc.priorities.Remove(item.priorityEntity)
		if removedKey != key {
			panic(errors.New("Wrong key in priority list"))
		}

		// Delete from map
		delete(lc.items, key)

		// Found and deleted
		return true
	}

	// Not found
	return false
}

func (lc *lruCacheImpl[Key, Value]) Length() int {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	return lc.len()
}

func (lc *lruCacheImpl[Key, Value]) Capacity() int {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	return lc.cap()
}

// Helper funcs

func (lc *lruCacheImpl[Key, Value]) len() int {
	return len(lc.items)
}

func (lc *lruCacheImpl[Key, Value]) cap() int {
	return lc.capacity
}

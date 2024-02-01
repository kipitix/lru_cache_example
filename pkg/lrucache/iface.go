package lrucache

type LRUCache[Key comparable, Value any] interface {
	// Add new value in cache
	// New value have highest priority
	// Returns true if action is successful
	// In case of key already exists - false will be returned
	// If the size is exceeded, the element with the lowest priority will be removed - true will be returned
	Add(key Key, value Value) bool

	// Returns the value under the key and the flag of its presence in the cache
	// If an element is in the cache, increases its priority
	Get(key Key) (Value, bool)

	// Removes an element from the cache, returns true if successful, false if the element is missing
	Remove(key Key) bool

	// Length of container
	Length() int

	// Capacity of container
	Capacity() int
}

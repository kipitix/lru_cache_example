# LRUCache

Самописный LRU cache

## Структура директорий

- application
	- Write business logic
- domain
	- Define interface
		- Repository interface for infrastructure
	- Define struct
		- Entity struct that represent mapping to data model
- infrastructure
	- Implements repository interface
	- Solves backend technical topics
		- e.x. message queue, persistence with RDB
- interfaces
	- Write HTTP handler and middleware

## LRU Cache

Расположен в директории pkg/lrucache

Сами элементы хранятся в контейнере map

Приоритеты хранятся в двусвязном списке

```go
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
```

## Для тестирования

```bash
make compose-up
make stress
```
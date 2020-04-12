package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type cacheItem struct {
	Key   Key
	Value interface{}
}

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	size     int
	queue    List
	elements map[Key]*listItem
}

func (cache *lruCache) isFull() bool {
	return cache.queue.Len() == cache.size
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	element, found := cache.elements[key]
	if found {
		element.value = cacheItem{
			Key:   key,
			Value: value,
		}
		cache.queue.MoveToFront(element)
		return true
	}

	if cache.isFull() {
		elmToDelete := cache.queue.Back()
		cache.queue.Remove(elmToDelete)
		delete(cache.elements, elmToDelete.value.(cacheItem).Key)
	}

	elm := cache.queue.PushFront(cacheItem{
		Key:   key,
		Value: value,
	})
	cache.elements[key] = elm
	return false
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	element, found := cache.elements[key]
	if found {
		cache.queue.MoveToFront(element)
		return element.value.(cacheItem).Value, true
	}
	return nil, false
}

func (cache *lruCache) Clear() {
	cache.queue = NewList()
	cache.elements = make(map[Key]*listItem, cache.size)
}

func NewCache(size int) Cache {
	return &lruCache{
		size:     size,
		queue:    NewList(),
		elements: make(map[Key]*listItem, size),
	}
}

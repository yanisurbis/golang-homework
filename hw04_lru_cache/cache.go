package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type cacheItem struct {
	// Place your code here
}

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Size int
	Queue List
	Elements map[Key]*listItem
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	element, found := cache.Elements[key]
	if found {
		cache.Queue.Remove(element)
		cache.Elements[key] = cache.Queue.PushFront(value)
		return true
	} else {
		if cache.Queue.Len() == cache.Size {
			// delete element
			elmToDelete := cache.Queue.Back()
			cache.Queue.Remove(elmToDelete)
			for key, elm := range cache.Elements {
				if elm == elmToDelete {
					delete(cache.Elements, key)
				}
			}
		}
		elm := cache.Queue.PushFront(value)
		cache.Elements[key] = elm
		return false
	}
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	element, found := cache.Elements[key]
	if found {
		cache.Queue.MoveToFront(element)
		return element.Value, true
	} else {
		return nil, false
	}
}

func (cache *lruCache) Clear() {
	cache.Queue = NewList()
	cache.Elements = make(map[Key]*listItem, cache.Size)
}

func NewCache(size int) Cache {
	return &lruCache{
		Size:     size,
		Queue:    NewList(),
		Elements: make(map[Key]*listItem, size),
	}
}
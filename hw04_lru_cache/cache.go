package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type cacheItem struct {
	Key Key
	Value interface{}
}

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Size     int
	Queue    List
	Elements map[Key]*listItem
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	element, found := cache.Elements[key]
	if found {
		cache.Queue.Remove(element)
		cache.Elements[key] = cache.Queue.PushFront(cacheItem{
			Key:   key,
			Value: value,
		})
		return true
	}

	if cache.Queue.Len() == cache.Size {
		elmToDelete := cache.Queue.Back()
		cache.Queue.Remove(elmToDelete)
		delete(cache.Elements, elmToDelete.Value.(cacheItem).Key)
	}

	elm := cache.Queue.PushFront(cacheItem{
		Key:   key,
		Value: value,
	})
	cache.Elements[key] = elm
	return false
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	element, found := cache.Elements[key]
	if found {
		cache.Queue.MoveToFront(element)
		return element.Value.(cacheItem).Value, true
	}
	return nil, false
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

package hw04_lru_cache //nolint:golint,stylecheck

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		cache := NewCache(2) // _ _
		cache.Set("1", 1) // 1 _
		cache.Set("2", 2) // 2 1
		cache.Set("3", 3) // 3 2

		_, found := cache.Get("1")
		require.False(t, found)

		v, found := cache.Get("2") // 2 3
		require.True(t, found)
		require.Equal(t, 2, v)

		v, found = cache.Get("3") // 3 2
		require.True(t, found)
		require.Equal(t, 3, v)

		cache.Set("4", 4) // 4 3
		_, found = cache.Get("2")
		require.False(t, found)

		cache.Set("5", 5) // 5 4
		_, found = cache.Get("3")
		require.False(t, found)

		cache.Set("4", 4) // 4 5
		cache.Set("6", 6) // 4 6
		_, found = cache.Get("5")
		require.False(t, found)

		cache.Clear() // _ _
		_, found = cache.Get("3")
		require.False(t, found)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove if task with asterisk completed

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}

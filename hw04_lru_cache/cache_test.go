package hw04lrucache

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
		// Write me
	})
}

func TestCacheMultithreading(t *testing.T) {
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

func TestOverload(t *testing.T) {
	c := NewCache(3)
	requiredKey := "first"
	c.Set(Key(requiredKey), 1)
	c.Set("second", 2)
	c.Set("third", 3)
	c.Set("fourth", 4)
	_, ok := c.Get(Key(requiredKey))
	require.False(t, ok)
}

func TestOverloadRemoveOld(t *testing.T) {
	c := NewCache(3)
	firstKey := "first"
	secondKey := "second"
	thirdKey := "third"
	fourthKey := "fourth"

	c.Set(Key(firstKey), 1)
	c.Set(Key(secondKey), 2)
	c.Set(Key(thirdKey), 3)

	c.Set(Key(firstKey), 4)
	c.Get(Key(secondKey))

	c.Set(Key(fourthKey), 10)

	_, ok := c.Get(Key(thirdKey))
	_, ok1 := c.Get(Key(secondKey))
	_, ok2 := c.Get(Key(firstKey))
	_, ok3 := c.Get(Key(fourthKey))
	require.False(t, ok)
	require.True(t, ok1, ok2, ok3)
}

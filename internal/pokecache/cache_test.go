package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	interval := 5 * time.Second
	cache := NewCache(interval)

	key := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	val := []byte("test-pokemon-data")

	cache.Add(key, val)

	got, exists := cache.Get(key)
	if !exists {
		t.Errorf("expected to find key %s in cache", key)
		return
	}

	if !bytes.Equal(got, val) {
		t.Errorf("expected cached value to be %s, got %s", string(val), string(got))
	}
}

func TestReapLoop(t *testing.T) {
	// Set a very short interval for testing expiration
	interval := 5 * time.Millisecond
	cache := NewCache(interval)

	key := "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20"
	cache.Add(key, []byte("test-data"))

	// Wait longer than the interval to allow reapLoop to clean up the entry[cite: 4]
	time.Sleep(interval + 20*time.Millisecond)

	_, exists := cache.Get(key)
	if exists {
		t.Errorf("expected key %s to be reaped and removed from the cache", key)
	}
}

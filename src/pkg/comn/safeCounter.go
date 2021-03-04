package comn

import "sync"

type SafeCounter struct {
	V map[string]int
	m sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.m.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.V[key]++
	c.m.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.m.Lock()
	// Lock so only one goroutine at a time can access the map c.V.
	defer c.m.Unlock()
	return c.V[key]
}

func (c *SafeCounter) Dec(key string) {
	c.m.Lock()
	c.V[key]--
	c.m.Unlock()
}

func (c *SafeCounter) Set(key string, val int) {
	c.m.Lock()
	c.V[key] = 0
	c.m.Unlock()
}

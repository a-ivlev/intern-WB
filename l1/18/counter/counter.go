package counter

import "sync"

type Count struct {
	mu *sync.Mutex
	count uint64
}

func NewCount() *Count {
	return &Count{
		mu: &sync.Mutex{},
		count: 0,
	}
}

func(c *Count) Inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func(c *Count) Get() uint64 {
	c.mu.Lock()
	count := c.count
	c.mu.Unlock()
	return count
}
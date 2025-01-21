package main

import (
	"fmt"
	"sync"
)

type Count struct {
	cnt int
	mu  sync.Mutex
}

// broke deadlock by mutex
func (c *Count) incrementByOne(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		c.mu.Unlock()
	}()

	c.mu.Lock()
	c.cnt += 1

}

func main() {
	var wg sync.WaitGroup

	count := Count{cnt: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go count.incrementByOne(&wg)
	}

	wg.Wait()

	fmt.Println(count.cnt)
}

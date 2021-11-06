package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type counter struct {
	counter int64
	mu      *sync.Mutex
}

func (c *counter) inc() {
	//c.mu.Lock()
	//c.counter++
	//c.mu.Unlock()
	atomic.AddInt64(&c.counter, 1)

}

func (c *counter) getRes() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter

}
func main() {
	t := time.Now()
	c := counter{
		mu: new(sync.Mutex),
	}
	wg := &sync.WaitGroup{}

	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			c.inc()
		}()
	}
	wg.Wait()
	fmt.Println(c.getRes())
	fmt.Println(time.Since(t))
}

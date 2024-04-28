package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type atomCounter struct {
	value int64
}

func (c *atomCounter) AddWithLock(i int64) {
	atomic.AddInt64(&c.value, i)
}

func (c *atomCounter) AddWithNonLock(i int64) {
	c.value += i
}

func (c *atomCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	locking()
	nonLocking()
}

func locking() {
	x := 100
	y := 4

	st := time.Now()

	fmt.Printf("start\n")

	var wg sync.WaitGroup
	counter := atomCounter{}

	for i := 0; i < x; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// =+4
			for j := 0; j < y; j++ {
				counter.AddWithLock(1)
			}
		}()
	}

	wg.Wait()

	fmt.Printf("%d ns\n", time.Since(st).Nanoseconds())

	// expected: 4 * 100 = 400
	log.Println("counter:", counter.Value())
}

func nonLocking() {
	x := 100
	y := 4

	st := time.Now()

	fmt.Printf("start\n")
	counter := atomCounter{}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			counter.AddWithLock(1)
		}
	}

	fmt.Printf("%d ns\n", time.Since(st).Nanoseconds())

	// expected: 4 * 100 = 400
	log.Println("counter:", counter.Value())
}

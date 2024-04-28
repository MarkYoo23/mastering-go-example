package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	count := 10
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func(x int) {
			if x != 0 {
				defer wg.Done()
			}

			fmt.Printf("%d ", x)
		}(i)
	}

	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Hour)
	}()

	wg.Wait()
}

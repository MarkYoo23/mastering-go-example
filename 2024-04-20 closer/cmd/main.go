package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 20; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}

	time.Sleep(1 * time.Second)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	go func(x int) {
		fmt.Printf("x: %d\n", x)
	}(10)

	time.Sleep(1 * time.Second)
}

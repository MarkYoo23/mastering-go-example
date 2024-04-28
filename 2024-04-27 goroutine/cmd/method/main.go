package main

import (
	"fmt"
	"time"
)

func main() {
	go printNumber(10)
	time.Sleep(1 * time.Second)
}

func printNumber(x int) {
	fmt.Printf("x: %d\n", x)
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go func(c chan int) {
		defer wg.Done()

		writeToCh(c, 10)
		fmt.Println("Exiting goroutine")
	}(c)

	fmt.Println("Read:", <-c)

	// 채널이 닫혔는지 확인
	value, ok := <-c
	if ok {
		fmt.Println("Channel is open")
	} else {
		fmt.Println("Channel is closed")
	}

	wg.Wait()

	fmt.Printf("ch value is %d\n", value)

}

func writeToCh(c chan int, x int) {
	c <- x
	close(c)
}

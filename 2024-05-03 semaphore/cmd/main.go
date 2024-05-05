package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

var workers = 4
var sem = semaphore.NewWeighted(int64(workers))

func main() {
	nJob := 6
	result := make([]int, nJob)

	ctx := context.TODO()

	for i := range result {
		err := sem.Acquire(ctx, 1)
		if err != nil {
			fmt.Println("Failed to acquire semaphore:", err)
			break
		}

		go func(i int) {
			defer sem.Release(1)
			temp := worker(i)
			result[i] = temp
		}(i)
	}

	for k, v := range result {
		fmt.Printf("result[%d] = %d\n", k, v)
	}
}

func worker(n int) int {
	square := n * n
	time.Sleep(1 * time.Second)
	return square
}

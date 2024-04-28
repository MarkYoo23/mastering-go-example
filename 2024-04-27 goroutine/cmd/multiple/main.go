package main

import "fmt"

func main() {
	count := 10

	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}

}

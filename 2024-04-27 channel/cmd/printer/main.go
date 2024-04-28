package main

import "fmt"

func main() {
	var ch = make(chan bool)

	for i := 0; i < 5; i++ {
		go printer(ch)
	}

	n := 0

	for i := range ch {
		println(i)

		if i == true {
			n++
		}

		if n > 2 {
			fmt.Println("n:", n)
			close(ch)
			break
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

func printer(ch chan bool) {
	ch <- true
}

package main

import "fmt"

func main() {
	willClose := make(chan complex64, 10)

	// 채널에 2개의 값을 발송후 닫음
	willClose <- -1
	willClose <- 1i
	close(willClose)

	// 채널을 닫았기 때문에 더 이상 값을 발송할 수 없음, 다만 값을 읽을 수는 있음
	read := <-willClose
	fmt.Println(read)
}

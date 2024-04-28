package main

import "sync/atomic"

func main() {
	var x int64 = 1
	atomic.AddInt64(&x, 1)
}

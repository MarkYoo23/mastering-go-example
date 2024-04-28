package main

import (
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	Method1()
}

func Method1() {
	var x int64 = 1
	atomic.AddInt64(&x, 1)
}

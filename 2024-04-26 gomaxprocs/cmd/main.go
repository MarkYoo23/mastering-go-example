package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("컴파일리 : %v \n", runtime.Compiler)
	fmt.Printf("아키텍처 : %v \n", runtime.GOARCH)
	fmt.Printf("운영체제 : %v \n", runtime.GOOS)
	// 항상 동시에 실행할 수 있는 CPU의 최대 개수, 현재 CPU 개수는 6개지만 SMT를 사용하면 12개로 인식
	_ = runtime.GOMAXPROCS(99)
	fmt.Printf("스레드 최대 개수 : %v \n", runtime.GOMAXPROCS(0))
	fmt.Printf("CPU의 개수 : %v \n", runtime.NumCPU())
	fmt.Printf("고루틴의 개수 : %v \n", runtime.NumGoroutine())
	fmt.Printf("버전 : %v \n", runtime.Version())
}

package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	fmt.Println("arch", runtime.GOARCH)
	fmt.Println("int", unsafe.Sizeof(int(0)))
	var x int
	_, err := fmt.Scan(&x)
	if err != nil {
		return
	}
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	if x < 10 && x > -10 {
		return x
	}
	origNum := x
	revNum := 0

	for origNum != 0 {
		revNum = revNum*10 + (origNum % 10)
		origNum /= 10
	}

	if revNum > 1<<31-1 || revNum < -(1<<31) {
		return 0
	}

	return revNum
}

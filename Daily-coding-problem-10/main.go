package main

import (
	"fmt"
	"time"
)

func scheduler(f func(a, b int) int, n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)

	fmt.Println(f(10,20))
}

func main() {
	f := func(a, b int) int {
		return a + b
	}
	scheduler(f, 1000)
}

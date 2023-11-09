package main

import (
	"fmt"
	"time"
)

func concurrentFib(n int) {
	ch := make(chan int)
	go fibonacci(n, ch)

	for v := range ch {
		fmt.Println(v)
	}
}

// TEST SUITE - Don't touch below this line

func test71(n int) {
	fmt.Printf("Printing %v numbers...\n", n)
	concurrentFib(n)
	fmt.Println("==============================")
}

func main() {
	test71(10)
	test71(5)
	test71(20)
	test71(13)
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}

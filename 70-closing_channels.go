package main

import (
	"fmt"
	"time"
)

func countReports(numSentCh chan int) int {
	total := 0
	for {
		numSent, ok := <-numSentCh
		if !ok {
			break
		}
		total += numSent
	}
	return total
}

// TEST SUITE - Don't touch below this line

func test70(numBatches int) {
	numSentCh := make(chan int)
	go sendReports(numBatches, numSentCh)

	fmt.Println("Start counting...")
	numReports := countReports(numSentCh)
	fmt.Printf("%v reports sent!\n", numReports)
	fmt.Println("========================")
}

func main() {
	test70(3)
	test70(4)
	test70(5)
	test70(6)
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
		fmt.Printf("Sent batch of %v reports\n", numReports)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

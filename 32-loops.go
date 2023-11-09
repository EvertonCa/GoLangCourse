package main

import (
	"fmt"
)

// go does not have while loop. The for loop can be used as while loop when using only de condition.
// for CONDITION {
//     // doing things until CONDITION is met
// }

// you can also use continue or break

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	// sections of the for can be optional.
	// for INITIAL; ; AFTER {
	//	// do something forever
	// }
	for i := 0; i < numMessages; i++ {
		totalCost += 1 + (float64(i) * 0.01)
	}
	return totalCost
}

// don't edit below this line

func test14(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("===============================================================")
}

func main() {
	test14(10)
	test14(20)
	test14(30)
	test14(40)
	test14(50)
}

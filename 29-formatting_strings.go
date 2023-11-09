package main

import (
	"fmt"
)

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' can not be sent",
		cost,
		recipient,
	)
}

// don't edit below this line

func test11(cost float64, recipient string) {
	s := getSMSErrorString(cost, recipient)
	fmt.Println(s)
	fmt.Println("====================================")
}

func main() {
	test11(1.4, "+1 (435) 555 0923")
	test11(2.1, "+2 (702) 555 3452")
	test11(32.1, "+1 (801) 555 7456")
	test11(14.4, "+1 (234) 555 6545")
}

//   fmt.Printf("I am %f years old", 10.523)
//   // I am 10.523000 years old
//
//   // The ".2" rounds the number to 2 decimal places
//   fmt.Printf("I am %.2f years old", 10.523)
//   // I am 10.53 years old

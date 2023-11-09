package main

import "fmt"

// IN GO, VARIABLES ARE PASSED BY VALUE

func main() {
	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	return sendsSoFar + sendsToAdd
}

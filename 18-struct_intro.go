package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

// don't edit below this line

func test2(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("====================================")
}

func main() {
	test2(messageToSend{
		phoneNumber: 148255510981,
		message:     "Thanks for signing up",
	})
	test2(messageToSend{
		phoneNumber: 148255510982,
		message:     "Love to have you aboard!",
	})
	test2(messageToSend{
		phoneNumber: 148255510983,
		message:     "We're so excited to have you",
	})
}

package main

import (
	"fmt"
	"strings"
)

func removeProfanity52(message *string) {
	if message == nil {
		return
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")
	*message = messageVal
}

// don't touch below this line

func test52(messages []string) {
	for _, message := range messages {
		if message == "" {
			removeProfanity52(nil)
			fmt.Println("nil message detected")
		} else {
			removeProfanity52(&message)
			fmt.Println(message)
		}
	}
}

func main() {
	messages := []string{
		"well shoot, this is awful",
		"",
		"dang robots",
		"dang them to heck",
		"",
	}

	messages2 := []string{
		"well shoot",
		"",
		"Allan is going straight to heck",
		"dang... that's a tough break",
		"",
	}

	test52(messages)
	test52(messages2)

}

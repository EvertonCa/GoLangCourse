package main

import "fmt"

func getMessageCosts(messages []string) []float64 {
	// func make([]T, length of the slice, capacity of the array) []T
	// if capacity is not defined, will default to length
	messageCosts := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		cost := float64(len(messages[i])) * 0.01
		messageCosts[i] = cost
	}
	return messageCosts
}

// don't edit below this line

func test17(messages []string) {
	costs := getMessageCosts(messages)
	fmt.Println("Messages:")
	for i := 0; i < len(messages); i++ {
		fmt.Printf(" - %v\n", messages[i])
	}
	fmt.Println("Costs:")
	for i := 0; i < len(costs); i++ {
		fmt.Printf(" - %.2f\n", costs[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test17([]string{
		"Welcome to the movies!",
		"Enjoy your popcorn!",
		"Please don't talk during the movie!",
	})
	test17([]string{
		"I don't want to be here anymore",
		"Can we go home?",
		"I'm hungry",
		"I'm bored",
	})
	test17([]string{
		"Hello",
		"Hi",
		"Hey",
		"Hi there",
		"Hey there",
		"Hi there",
		"Hello there",
		"Hey there",
		"Hello there",
		"General Kenobi",
	})
}

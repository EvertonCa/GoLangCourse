package main

import "fmt"

func printReports(messages []string) {
	for _, message := range messages {
		printCostReport(func(m string) int {
			return len(m) * 2
		}, message)
	}
}

// don't touch below this line

func test49(messages []string) {
	defer fmt.Println("====================================")
	printReports(messages)
}

func main() {
	test49([]string{
		"Here's Johnny!",
		"Go ahead, make my day",
		"You had me at hello",
		"There's no place like home",
	})

	test49([]string{
		"Hello, my name is Inigo Montoya. You killed my father. Prepare to die.",
		"May the Force be with you.",
		"Show me the money!",
		"Go ahead, make my day.",
	})
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}

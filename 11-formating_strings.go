package main

import "fmt"

func main() {
	const name = "Saul Goodman"
	const openRate = 30.5

	// fmt.Printf - Prints a formatted string to standard out.
	// fmt.Sprintf() - Returns the formatted string
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)

	fmt.Println(msg)

	// ------------- Interpolate the default representation
	fmt.Printf("I am %v years old", 10)
	// I am 10 years old

	fmt.Printf("I am %v years old", "way too many")
	// I am way too many years old

	// ------------- Interpolate a string
	fmt.Printf("I am %s years old", "way too many")
	// I am way too many years old

	// ------------- Interpolate an integer in decimal form
	fmt.Printf("I am %d years old", 10)
	// I am 10 years old

	// ------------- Interpolate a decimal
	fmt.Printf("I am %f years old", 10.523)
	// I am 10.523000 years old

	// The ".2" rounds the number to 2 decimal places
	fmt.Printf("I am %.2f years old", 10.523)
	// I am 10.53 years old

}

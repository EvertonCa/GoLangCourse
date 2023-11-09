package main

import "fmt"

func main() {
	const secondsInMinute = 60
	const minutesInHour = 60

	// constants will ONLY be calculated at compile time, not runtime!
	const secondsInHour = secondsInMinute * minutesInHour

	// don't edit below this line
	fmt.Println("number of seconds in an hour:", secondsInHour)
}

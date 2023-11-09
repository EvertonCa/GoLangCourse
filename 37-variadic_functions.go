package main

import "fmt"

// ...int is just a slice of integers. []int
func sum(nums ...float64) float64 {
	sum := 0.0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		sum += num
	}
	return sum
}

// don't edit below this line

func test18(nums ...float64) {
	// spread operator = pass a slice as variadic
	total := sum(nums...)
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %.2f\n", total)
	fmt.Println("===== END REPORT =====")
}

func main() {
	test18(1.0, 2.0, 3.0)
	test18(1.0, 2.0, 3.0, 4.0, 5.0)
	test18(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
	test18(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0)
}

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// don't touch below this line

type emailBill struct {
	costInPennies int
}

func test48(bills []emailBill) {
	defer fmt.Println("====================================")
	countAdder, costAdder := adder(), adder()
	for _, bill := range bills {
		fmt.Printf("You've sent %d emails and it has cost you %d cents\n", countAdder(1), costAdder(bill.costInPennies))
	}
}

func main() {
	test48([]emailBill{
		{45},
		{32},
		{43},
		{12},
		{34},
		{54},
	})

	test48([]emailBill{
		{12},
		{12},
		{976},
		{12},
		{543},
	})

	test48([]emailBill{
		{743},
		{13},
		{8},
	})
}

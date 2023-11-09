package main

import (
	"errors"
	"fmt"
)

func divide2(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("no dividing by 0")
	}
	return x / y, nil
}

// don't edit below this line

func test13(x, y float64) {
	defer fmt.Println("====================================")
	fmt.Printf("Dividing %.2f by %.2f ...\n", x, y)
	quotient, err := divide2(x, y)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient: %.2f\n", quotient)
}

func main() {
	test13(10, 0)
	test13(10, 2)
	test13(15, 30)
	test13(6, 3)
}

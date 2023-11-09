package main

import (
	"fmt"
)

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (e email) print() {
	fmt.Println(e.body)
}

// don't touch below this line

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
}

func print(p printer) {
	p.print()
}

func test8(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("====================================")
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "hello there",
	}
	test8(e, e)
	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	test8(e, e)
	e = email{
		isSubscribed: true,
		body:         "Are you free for a chat?",
	}
	test8(e, e)
	e = email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
	}
	test8(e, e)
}
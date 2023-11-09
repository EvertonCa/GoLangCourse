package main

import (
	"fmt"
)

func getExpenseReport(e expense2) (string, float64) {
	// "em" is a new email2 cast from "e"
	// which is an instance of an expense2.
	// "ok" is a bool that is true if s was an email2
	// or false if s isn't an email2
	em, ok := e.(email2)
	if ok {
		return em.toAddress, em.cost()
	}

	sm, ok := e.(sms)
	if ok {
		return sm.toPhoneNumber, sm.cost()
	}

	return "", 0.0
}

// don't touch below this line

func (e email2) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

type expense2 interface {
	cost() float64
}

type email2 struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func estimateYearlyCost(e expense2, averageMessagesPerYear int) float64 {
	return e.cost() * float64(averageMessagesPerYear)
}

func test9(e expense2) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case email2:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	case sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("====================================")
	}
}

func main() {
	test9(email2{
		isSubscribed: true,
		body:         "hello there",
		toAddress:    "john@does.com",
	})
	test9(email2{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "jane@doe.com",
	})
	test9(email2{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "elon@doe.com",
	})
	test9(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555509832",
	})
	test9(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555504444",
	})
	test9(invalid{})
}

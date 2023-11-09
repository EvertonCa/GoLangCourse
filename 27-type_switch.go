package main

import (
	"fmt"
)

func getExpenseReport2(e expense3) (string, float64) {
	switch v := e.(type) {
	case email3:
		return v.toAddress, v.cost()
	case sms3:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
}

// don't touch below this line

func (e email3) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms3) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid3) cost() float64 {
	return 0.0
}

type expense3 interface {
	cost() float64
}

type email3 struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms3 struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid3 struct{}

func estimateYearlyCost3(e expense2, averageMessagesPerYear int) float64 {
	return e.cost() * float64(averageMessagesPerYear)
}

func test16(e expense2) {
	address, cost := getExpenseReport2(e)
	switch e.(type) {
	case email3:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	case sms3:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("====================================")
	}
}

func main() {
	test16(email3{
		isSubscribed: true,
		body:         "hello there",
		toAddress:    "john@does.com",
	})
	test16(email3{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "jane@doe.com",
	})
	test16(email3{
		isSubscribed: false,
		body:         "Wanna catch up later?",
		toAddress:    "elon@doe.com",
	})
	test16(sms3{
		isSubscribed:  false,
		body:          "I'm a Nigerian prince, please send me your bank info so I can deposit $1000 dollars",
		toPhoneNumber: "+155555509832",
	})
	test16(sms3{
		isSubscribed:  false,
		body:          "I don't need this",
		toPhoneNumber: "+155555504444",
	})
	test16(invalid3{})
}

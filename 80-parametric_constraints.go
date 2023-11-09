package main

import (
	"fmt"
)

type biller[C customer] interface {
	Charge(C) bill
	Name() string
}

// don't edit below this line

type userBiller struct {
	Plan string
}

func (ub userBiller) Charge(u user80) bill {
	amount := 50.0
	if ub.Plan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

func (sb userBiller) Name() string {
	return fmt.Sprintf("%s user biller", sb.Plan)
}

type orgBiller struct {
	Plan string
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.0
	}
	return bill{
		Customer: o,
		Amount:   amount,
	}
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type user80 struct {
	UserEmail string
}

func (u user80) GetBillingEmail() string {
	return u.UserEmail
}

type org struct {
	Admin user80
	Name  string
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}

func main() {
	testBiller[user80](
		userBiller{Plan: "basic"},
		user80{UserEmail: "joe@example.com"},
	)
	testBiller[user80](
		userBiller{Plan: "basic"},
		user80{UserEmail: "samuel.boggs@example.com"},
	)
	testBiller[user80](
		userBiller{Plan: "pro"},
		user80{UserEmail: "jade.row@example.com"},
	)
	testBiller[org](
		orgBiller{Plan: "basic"},
		org{Admin: user80{UserEmail: "challis.rane@example.com"}},
	)
	testBiller[org](
		orgBiller{Plan: "pro"},
		org{Admin: user80{UserEmail: "challis.rane@example.com"}},
	)
}

func testBiller[C customer](b biller[C], c C) {
	fmt.Printf("Using '%s' to create a bill for '%s'\n", b.Name(), c.GetBillingEmail())
	bill := b.Charge(c)
	fmt.Printf("Bill created for %v dollars\n", bill.Amount)
	fmt.Println(" --- ")
}

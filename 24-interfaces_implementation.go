package main

import (
	"fmt"
)

// interfaces are implemented implicitly.

// args can have names
//
//	type Copier interface {
//		Copy(sourceFile string, destinationFile string) (bytesCopied int)
//	}
type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

// don't touch below this line

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func test7(e employee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("====================================")
}

func main() {
	test7(fullTime{
		name:   "Jack",
		salary: 50000,
	})
	test7(contractor{
		name:         "Bob",
		hourlyPay:    100,
		hoursPerYear: 73,
	})
	test7(contractor{
		name:         "Jill",
		hourlyPay:    872,
		hoursPerYear: 982,
	})
}

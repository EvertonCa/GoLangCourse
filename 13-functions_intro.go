package main

import "fmt"

// you can only declare the type of the last args, if they are of the same type
// func concat(s1, s2 string) string {
func concat(s1 string, s2 string) string {
	return s1 + s2
}

func main() {
	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}

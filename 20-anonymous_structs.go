package main

import "fmt"

func main() {
	myCar := struct {
		Make  string
		Model string
	}{
		Make:  "tesla",
		Model: "model 3"}

	fmt.Println(myCar.Make)
}

package main

import "fmt"

type sender struct {
	rateLimit int
	// "user2" is embedded, so the definition of a "sender" now also additionally contains all the fields of the user2 struct
	user2
}

type user2 struct {
	name   string
	number int
}

// don't edit below this line

func test4(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("====================================")
}

func main() {
	test4(sender{
		rateLimit: 10000,
		user2: user2{
			name:   "Deborah",
			number: 18055558790,
		},
	})
	test4(sender{
		rateLimit: 5000,
		user2: user2{
			name:   "Sarah",
			number: 19055558790,
		},
	})
	test4(sender{
		rateLimit: 1000,
		user2: user2{
			name:   "Sally",
			number: 19055558790,
		},
	})
}

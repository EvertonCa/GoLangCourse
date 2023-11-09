package main

import "fmt"

type Message struct {
	Recipient string
	Text      string
}

// Don't touch above this line

func sendMessage50(m Message) {
	fmt.Printf("To: %v\n", m.Recipient)
	fmt.Printf("Message: %v\n", m.Text)
}

// Don't touch below this line

func test50(recipient string, text string) {
	m := Message{Recipient: recipient, Text: text}
	sendMessage50(m)
	fmt.Println("=====================================")
}

func main() {
	test50("Lane", "Textio is getting better everyday!")
	test50("Allan", "This pointer stuff is weird...")
	test50("Tiffany", "What time will you be home for dinner?")
}

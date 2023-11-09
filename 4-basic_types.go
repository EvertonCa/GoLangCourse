package main

import "fmt"

func main() {
	// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64 and uintptr
	// byte = alias for uint8
	// rune = alias for int32 - represents a Unicode code point
	var smsSendingLimit int

	// float32, float64, complex64 and complex128
	var costPerSMS float64

	var hasPermission bool
	var username string
	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}

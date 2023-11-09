package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// has a receiver of authenticationInfo struct
func (authInfo authenticationInfo) getBasicAuth() string {
	return "Authorization: Basic " + authInfo.username + ":" + authInfo.password
}

// don't touch below this line

func test5(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}

func main() {
	test5(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test5(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test5(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}

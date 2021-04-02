package main

import (
	"fmt"
	"os"
)

var users = map[string]string{
"305": "Sue",
"204": "Bob",
"631": "Jake",
"073": "Tracy",
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No arguments given.")
		os.Exit(1)
	}

	if users[os.Args[1]] == "" {
		fmt.Println("Non existing user")
		os.Exit(1)
	}
	fmt.Println("Hi,", users[os.Args[1]])
}
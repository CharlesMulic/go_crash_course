package main

import "fmt"

func main() {
	var emails map[string]string = make(map[string]string)
	// emails := make(map[string]string) // defines a map of string keys to string objects
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	// declaration and initialization
	// emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	delete(emails, "Bob")
	fmt.Println(emails)
	fmt.Println(len(emails))
}

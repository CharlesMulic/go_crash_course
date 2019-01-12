package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string

	firstName, lastName, city, gender string // can decalre all of the same type on one line
	age                               int
}

// value receiver, doesn't alter
func (person Person) greet() string { // in this case "person" is similar to "this"
	// person.firstName = person.firstName + "!" // why does this work for a value receiver?
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age)
}

// pointer receiver
func (person *Person) hasBirthday() {
	person.age++
}

func (person *Person) getMarried(surname string) {
	if person.gender == "M" || person.gender == "m" {
		return
	}
	person.lastName = surname
}

func main() {
	person1 := Person{
		firstName: "Charlie",
		lastName:  "Mulic",
		city:      "Minneapolis",
		gender:    "M",
		age:       33, // why does there need to be a comma here?
	}

	person2 := Person{"Allison", "Something", "Minneapolis", "F", 29} // also works without naming variables, order/type matters

	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)

	person1.hasBirthday()
	fmt.Println(person1.greet())

	fmt.Println(person2.greet())
	person2.getMarried("Test")
	fmt.Println(person2.greet())
}

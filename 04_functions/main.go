package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

// func getSum(num1 int, num2 int) int {
func getSum(num1, num2 int) int { // shorthand for when both parameters are the same type
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Charlie"))
	fmt.Println(getSum(20, 7))
}

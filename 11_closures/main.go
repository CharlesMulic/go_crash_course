package main

import "fmt"

// a function returning a function/closure
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// sum := adder()

	sum := 0
	calc := func(x int) int { // inline anonymous function/closure
		sum += x
		return sum
	}

	for i := 0; i < 10; i++ {
		// fmt.Println(sum(i))
		fmt.Println(calc(i))
	}
}

package main

import "fmt"

func main() {
	// arrays have to be a fixed length and known types
	// slices are basically arrays without fixed length

	// var fruitArr [2]string

	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// fruitArr := [2]string{"Apple", "Orange"} // declare and init shorthand
	fruitArr := []string{"Apple", "Orange", "Grape", "Cherry"} // no size declared, this is a slice (like array list?)

	fmt.Println(fruitArr)
	fmt.Println(len(fruitArr)) // get length
	fmt.Println(fruitArr[1:3]) // get a range
	fmt.Println(fruitArr[1])
}

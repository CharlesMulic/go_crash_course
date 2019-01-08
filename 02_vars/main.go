package main

import "fmt"

func main() {
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// using var
	// var name string = "Charlie"
	// var name = "Charlie" // type is inferred from assignment, don't need to add string as in above
	// name := "Charlie"
	name, email := "Charlie", "charles.mulic@gmail.com" // multiple declarations and assignments in one line

	var age = 33 // inferred int
	// var age int32 = 33

	const isCool = true
	// isCool = false

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", name) // %T gets type
}

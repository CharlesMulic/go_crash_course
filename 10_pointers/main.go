package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a)
	fmt.Println(b)        // b is a pointer, not the value
	fmt.Printf("%T\n", b) // get the type, is *int (pointer to int)
	fmt.Println(*b)       // dereference the pointer to get the value

	a++ // incrementing a will also increment b since b is just a pointer to a
	fmt.Println(a)
	fmt.Println(*b) // equivalent to *&a (dereference a pointer to a, same as just a)

	*b = 10
	fmt.Println(a, *b)

	c := []int{1, 2, 3, 4, 5}

	fmt.Println(c[1])
	// fmt.Println(*(&c[1] + 4)) // there is no pointer arithmetic in Go, would have to use "unsafe" package
}

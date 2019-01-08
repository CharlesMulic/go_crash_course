package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x < y { // convention to not use parentheses, but you can just fine
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is not less than %d\n", x, y)
	}

	color := "blue"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is not blue or red")
	}

	switch color {
	case "red":
		fmt.Print("color is red")
	case "blue":
		fmt.Print("color is blue")
	default:
		fmt.Print("color is not blue or red")
	}
}

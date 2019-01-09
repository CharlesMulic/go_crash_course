package main

import "fmt"

func main() {
	ids := []int{33, 25, 32, 36, 56}

	for i, id := range ids { // range seems to (at least for an array/slice) give you an iterator with index and value for each element
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids { // use _ for unused values?
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Maps
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for key, value := range emails {
		fmt.Printf("%s: %s\n", key, value)
	}

	for key := range emails {
		fmt.Println("Name: " + key)
	}
}

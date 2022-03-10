package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world.")
	testContainsDups()
}
func testContainsDups() {
	a := []int{1, 2, 3, 1}
	b := []int{1, 2, 3, 4}
	c := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	fmt.Printf("A: %t\tB: %t\tC: %t", containsDuplicate(a), containsDuplicate(b), containsDuplicate(c))

}

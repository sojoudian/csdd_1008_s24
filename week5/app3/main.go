package main

import "fmt"

func main() {
	a := 2
	b := 3
	addNumbers(a, b)

}

func addNumbers(a int, b int) int {
	var c int = a + b
	fmt.Println("Sum of a and b is", c)
	return c
}

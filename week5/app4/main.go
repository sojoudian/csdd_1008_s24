package main

import (
	"fmt"
)

func main() {
	testNumber := 90
	checkEvenOdd((testNumber))
}

func checkEvenOdd(num int) {
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

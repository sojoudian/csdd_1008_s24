package main

import "fmt"

func main() {
	name := "Maziar"
	message := "Hello"
	age := 35
	greeting(message, name, age)
}

func greeting(message string, name string, age int) {
	fmt.Println(message, name, ". You are ", age, "old")
}

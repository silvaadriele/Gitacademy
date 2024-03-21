package main

import "fmt"

func main() {
	var number1 int
	var number2 int
	var sum int
	fmt.Println("Enter number 1")
	fmt.Scanln(&number1)
	fmt.Println("Enter number 2")
	fmt.Scanln(&number2)

	sum = number1 + number2
	fmt.Printf("Result %d", sum)

}

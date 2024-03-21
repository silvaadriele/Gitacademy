package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter the number")
	fmt.Scanln(&number)

	if number < 0 {
		fmt.Println("<0")
	} else if number == 0 {
		fmt.Println("0")
	} else {
		fmt.Println(">0")
	}
}

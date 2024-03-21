package main

import "fmt"

func main() {
	fmt.Println("Enter you name")
	var name string

	fmt.Scanln(&name)

	fmt.Printf("Hello %s", name)
}

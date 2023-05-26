package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numberString string

	fmt.Print("Enter a number: ")
	fmt.Scanln(&numberString)
	num, err := strconv.Atoi(numberString)

	if err != nil {
		fmt.Println("Invalid input!")
		return
	}

	if num%2 == 0 {
		fmt.Println("The given number is even!")
	} else {
		fmt.Println("The given number is odd!")
	}
}

package main

import "fmt"

func main() {
	num := 34
	fact := 1
	for i := 1; i <= num; i++ {
		fact *= i
	}
	fmt.Println("Factorial of ", num, "is ", fact)
}

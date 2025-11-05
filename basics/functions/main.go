package main

import "fmt"

func main() {
	fmt.Println("Functions Demo")
	simpleOutput()
	additionsans := additionOfTwoNumbers(12, 12)
	fmt.Println("Addition ans", additionsans)
	ans := multiplicationOfNumbers(12, 12)
	fmt.Println("Multiplication ans", ans)
}

func simpleOutput() {
	fmt.Println("Simple function with no return type and only prints the output")
}

func additionOfTwoNumbers(a, b int) int {
	return a + b
}

func multiplicationOfNumbers(a, b int) (result int) {
	result = a * b
	return
}

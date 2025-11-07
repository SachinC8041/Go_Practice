package main

import "fmt"

func main() {
	fmt.Println("we are going to learn Go language in detail now. Are you ready guys ???")

	var name string = "Sachin Tendulkar"
	fmt.Println("The name is :", name)

	var age int = 23
	fmt.Println("The age is:", age)

	var temperature float64 = 23.2323
	fmt.Println("The temperature is :", temperature)

	var isApplicable bool = true
	if isApplicable {
		fmt.Println("The candidate is Applicable")
	} else {
		fmt.Println("The candidate is Not Applicable")
	}
	fmt.Printf("The name of the person is %s and his age is %d who is standing in place where temperature is %.2f\n", name, age, temperature)
}

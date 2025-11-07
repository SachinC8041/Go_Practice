package main

import "fmt"

func main() {
	fmt.Println("Condition statements and Operators lecture")
	var count int
	fmt.Println("Please enter your count value")
	fmt.Scanln(&count)
	if count < 0 {
		fmt.Println("Invalid INput")
	} else if count == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Final Count", 100*count)
	}

	var x, y, z int
	fmt.Println("Please enter your 3 integer input value")
	fmt.Scanln(&x, &y, &z)
	if x > y && y <= z {
		fmt.Printf("The value %d is greater than other 2 value\n", x)
	} else if y > x && y > z {
		fmt.Printf("The value %d is greater than other 2 values \n", y)
	} else {
		fmt.Println("The highest value is :", z)
	}

	// && , || , ! are the operators for AND , OR , NOT
	// < , <= , > , >= , == are the operators for GREATER THAN , LESS THAN , EQUALS

	fmt.Println("Please enter your input for switch case")
	var rollNo int
	fmt.Scanln(&rollNo)
	switch rollNo {
	case 10:
		fmt.Println("Input 1")
	case 20:
		fmt.Println("Input 2")
	case 30:
		fmt.Println("Input 3")
	default:
		fmt.Println("Invalid Input")
	}

}

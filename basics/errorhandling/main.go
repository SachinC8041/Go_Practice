package main

import "fmt"

func main() {
	fmt.Println("We are going to learn Error Handling and underscore operator")
	divisionans, error := division(12, 0)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(divisionans)
	}

}

func division(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("Divisor is zero")
	}
	return a / b, nil
}

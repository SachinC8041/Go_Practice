package main

import "fmt"

func main() {
	fmt.Println("Loops lecture in Go")

	input := 12
	for i := 1; i <= 10; i++ {
		fmt.Println(input * i)
	}

	var input2 = [4]int{12, 13, 14, 15}
	for i, v := range input2 {
		fmt.Printf("Index %d Value %d\n", i, v)
	}
}

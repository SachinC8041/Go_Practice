package main

import "fmt"

func main() {
	fmt.Println("Slice lecture")

	//Method 1 to to make slices
	numbers := []int{12, 323, 23515}
	fmt.Println(numbers)
	var names []string
	names = append(names, "Saurav", "Srikant", "Virendra", "sachin", "Rahul", "Balaji")
	fmt.Println(names)

	// Method 2 to make slices using make method
	animals := make([]int, 5, 10)
	fmt.Println(animals)
	fmt.Println(len(animals))
	fmt.Println(cap(animals))

	//subslices
	animalSubslices := names[:3]
	fmt.Println(animalSubslices)

}

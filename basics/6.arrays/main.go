package main

import "fmt"

func main() {
	fmt.Println("Arrays lecture")
	// Method 1 to declare Arrays
	var names [3]string
	names[0] = "Sachin"
	names[1] = "Virat Kohli" //filling the array elements
	names[2] = "Gautam"
	fmt.Println("names of the sportpersons are :", names)
	names[2] = "Gautam Gambhir" //modifying the array elements

	// Method 2 to declare Arrays and Initialize it
	var numbers = [5]int{11, 22, 33, 44, 55}
	fmt.Println("The numbers array is :", numbers)

	// Iterate through arrays
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index", i, ":", numbers[i])
	}

	// Iterate through arrays : range method
	for i, v := range numbers {
		fmt.Printf("Index %d Value %d \n", i, v)
	}

	//Zero value demonstration of arrays
	var integer [4]int
	var stringvalue [4]string
	var boolean [3]bool
	var floatdemonstration [3]float32
	fmt.Println("Integer zero value", integer)
	fmt.Println("String zero value", stringvalue) //for string value demonstration you can also use Quoted string i.e. %Q
	fmt.Println("Boolean zero value", boolean)
	fmt.Println("Float zero value", floatdemonstration)

	//2D arrays demonstration of arrays
	var stuMarks = [3][3]int{{12, 13, 12}, {14, 12, 16}, {11, 12, 9}}
	fmt.Println(stuMarks)

}

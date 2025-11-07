package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("We are going to learn about How to take input from User in Go")
	var rollNO int
	fmt.Println("Please enter your Roll No")
	fmt.Scanln(&rollNO)
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("name is:", name)

}

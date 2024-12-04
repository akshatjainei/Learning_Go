package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the Pizzeria"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating of our Pizza")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
}

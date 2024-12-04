package main

import (
	"fmt"
	"slices"
)

func main() {
	fruitList := []string{"Apple", "Melon"}
	fmt.Println(fruitList)

	// Example of using slices package
	slices.Sort(fruitList)
	fmt.Println("Sorted:", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	scores := make([]int, 3)
	scores[0] = 1
	scores[1] = 2
	scores[2] = 3

	fmt.Println(scores)
}

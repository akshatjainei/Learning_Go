package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello")
	content := "This needs to be writen"

	file, err := os.Create("./mytxt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println(length)
	defer file.Close()
}

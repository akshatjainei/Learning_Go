package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {
	newUser := []User{
		{"Akshat", 23, "Meta", "hello", []string{"nik", "vik"}},
		{"Shanks", 21, "Tel", "Mel", []string{"Bing", "Ding "}},
		{"Racos", 23, "wad", "Adhd", nil},
	}
	finalJson, err := json.Marshal(newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(finalJson)
}

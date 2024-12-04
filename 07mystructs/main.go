package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    int16
	Status bool
}

func main() {
	akshat := User{"Akshat", "jainakshat447@gmail.com", 20, true}
	fmt.Println(akshat.Name)

	days := [7]string{"Mon", "Tue", "wed", "Thurs", "Fri", "Sat", "Sun"}

	for i := range days {
		fmt.Println(days[i])
	}
}

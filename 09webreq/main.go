package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of responce is %T \n", res)

	defer res.Body.Close()

	dataBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	content := string(dataBytes)
	fmt.Println(content)
}

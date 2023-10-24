package main

import "fmt"

func main(){
	var testmap map[int]string

	println(testmap)

	test2map := make(map[int]string)

	println(test2map)

	tickers := map[string]string{
		"goog": "Google Coperation",
		"ho": "hello",
	}

	fmt.Println(tickers)

}
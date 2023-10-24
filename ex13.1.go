package main

import "fmt"

type House struct {
	Address string
	Size int
	Price float64
	Type string
}

func main() {
	var house House

	house.Address = "서울시 강동구 ..."
	house.Size = 29
	house.Price = 9.8
	house.Type = "아파트"

	fmt.Println(house)

	house = House{Address: "서울시 강남구", Price: 21}

	fmt.Println(house)
}
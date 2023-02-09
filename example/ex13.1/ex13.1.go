package main

import "fmt"

type House struct {
	Adddress string
	Size     int
	Price    float64
	Category string
}

func main() {
	var house House
	house.Adddress = "수원시 권선구"
	house.Size = 28
	house.Price = 10
	house.Category = "APT"

	var house2 House = House{
		Size:     28,
		Category: "APT",
	}

	fmt.Println(house, house2)
}

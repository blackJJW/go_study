package main

import "fmt"

// 구조체는 여러 필드를 묶어서 사용하는 타입
type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func main() {
	var house House
	house.Address = "서울시 강남구 ..."
	house.Size = 29
	house.Price = 10
	house.Category = "apartment"

	fmt.Printf("%v\n", house)
}

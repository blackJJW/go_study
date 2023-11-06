package main

func main() {
	a := 3 // int
	var b float64 = 3.5

	var c int = b // Error - float64 변수를 int에 대입 불가
	d := a * b    // Error - 다른 타입인 int 변수와 float64 연산 불가

	var e int64 = 7
	f := a * e // Error - a는 int 타입, e는 int64 타입으로 같은 정수값이지만
	// 타입이 달라서 연산 불가

}

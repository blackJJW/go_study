package main

import "fmt"

func main() {
	var a int = 3 // 기본 타입
	var b int     // 초깃값 생략, 초깃값은 타입별 기본값으로 대체
	var c = 4     // 타입 생략, 변수 타입은 우변 값의 타입이 됨
	d := 5        // 선언 대입문 :=을 사용해서 var 키워드와 타입 생략

	var e = "Hello"
	f := 3.14

	fmt.Println(a, b, c, d)
	fmt.Println(e, f)
}

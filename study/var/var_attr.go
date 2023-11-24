package main

import "fmt"

func main() {
	// 변수의 4가지 속성
	// 1. 이름 2. 값 3. 주소 4. 타입
	// 이름 : 프로그래머용
	// 주소 : 컴퓨터
	// 타입 : 정수, 실수, 문자열 등, 사이즈는 고정
	//	int32 : 정수형 - 32bin(4byte)

	var a int16 = 3456 // a : int16 , 2byte
	var b int8 = int8(a)

	fmt.Println(a, b)

}

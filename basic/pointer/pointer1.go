package main

import "fmt"

// 포인터는 메모리 주로를 값으로 갖는 타입

func main() {
	var a int = 500
	var p *int

	p = &a

	fmt.Printf("p의 값 : %p\n", p)
	fmt.Printf("p가 가리키는 메모리의 값 : %d\n", *p)

	*p = 100

	fmt.Printf("a의 값 : %d\n", a)

}

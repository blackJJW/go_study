package main

import "fmt"

// 구체화된 객체가 아닌 추상화된 상호작용으로 관계를 표현

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("Hi! %d  %s", s.Age, s.Name)
}

func main() {
	student := Student{"abc", 12}
	var stringer Stringer

	stringer = student
	fmt.Printf("%s\n", stringer.String())
}

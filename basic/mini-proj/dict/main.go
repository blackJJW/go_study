package main

import (
	"fmt"

	"github.com/blackJJW/learngo/mini-proj/dict/mydict"
)

func main() {

	word := "hello"
	definition := "Greeting"

	dictionary := mydict.Dictionary{"first": "First word"}
	definition_2, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition_2)
	}

	err_1 := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err_1)
	}
	hello_1, err_1 := dictionary.Search(word)
	fmt.Println(hello_1)
}

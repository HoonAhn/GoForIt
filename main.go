package main

import (
	"fmt"
	"main/myDict"
)

func main() {
	dictionary := myDict.Dictionary{}
	newword := "hello"
	newdef := "World"
	err := dictionary.Add(newword, newdef)
	if err != nil {
		fmt.Println(err)
	}
	err2 := dictionary.Update(newword, "hoon")
	if err2 != nil {
		fmt.Println(err2)
	}
	dictionary.Delete("yo")
	definition2, _ := dictionary.Search(newword)
	fmt.Println(definition2)
}

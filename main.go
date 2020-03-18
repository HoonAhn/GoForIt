package main

import (
	"fmt"
)

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	fav := []string{"kimchi", "real ramen"}
	hoon := person{name: "bacon", favFood: fav, age: 20}
	fmt.Println(hoon)
}

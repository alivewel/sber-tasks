package main

import (
	"fmt"
)

// что выведет программа?

func main() {
	slice := make([]int, 0, 1000)
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)     // ???
	process(slice)
	fmt.Println(slice)     // ???
	fmt.Println(slice[:6]) // ???
}

func process(slice []int) {
	for index := range slice {
		slice[index] = 0
	}
}

// Что такое слайс? Почему он так называется?
// Как работает функция append?
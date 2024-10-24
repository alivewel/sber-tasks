package main

import (
	"fmt"
)

// что выведет программа?

func main() {
	slice := make([]int, 0, 1000)
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)     // 1 2 3
	process(slice)
	fmt.Println(slice)     // 3 нуля
	fmt.Println(slice[:6]) // 6 нулей (последний не включительно)
}

func process(slice []int) {
	for index := range slice { // 3 итерации цикла, len = 3
		fmt.Println(index)
		slice[index] = 0
	}
}

// Что такое слайс? Почему он так называется?
// см. README
// [1 2 3]
// [0 0 0]
// [0 0 0 0 0 0]

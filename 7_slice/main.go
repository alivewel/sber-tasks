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
	slice = append(slice, 4, 5, 6)
}

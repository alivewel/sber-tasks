package main

// что выведет программа?

func main() {
	for _, val := range []int{1, 2, 3} {
		go println(val)
	}
}

package main

// что выведет программа?

func main() {
	println(f())
}

func f() string {
	s := "Test"
	s[0] = 'R'
	return s
}

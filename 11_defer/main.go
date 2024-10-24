package main

// Что и почему напечатает программа?

func f1() int {
	x := 1
	defer func() {
		x += 1
	}()
	return x
}

func f2() (x int) {
	x = 1
	defer func() {
		x += 1
	}()
	return x
}

func main() {
	println(f1())
	println(f2())
}

// Что такое defer?
// defer - отложенное выполнение программы
// defer функции отрабатываются, после выхода из основной функции
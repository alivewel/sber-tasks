package main

// реализовать функцию, которая возвращает ошибку, не используя импорт других библиотек

type myError struct {}

func (e *myError) Error() string {
	return "my error"
}

func handle() error {
	return &myError{}
}

func main() {
	err := handle()
	if err != nil {
		println(err.Error())
	}
}

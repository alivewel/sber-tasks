package main

func main() {
	str := "Привет!"
	charCount := 0

	// посчитать кол-во символов в строке
	charCount = len([]rune(str)) // Преобразуем строку в срез рун и берем его длину

	if charCount == 7 {
		println("Success!")
	}
}

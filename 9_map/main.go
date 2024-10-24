package main

import (
	"math/rand"
)

func randomInt(max int) int {
	return rand.Intn(max) // Возвращает случайное число от 0 до max-1
}

func hash1() int {
	return randomInt(100)
}

func hash2() int {
	return 1
}

// 47:45
// какая функция для вычисления хеша лучше? Какие плюсы и минусы каждой функции?

// Коллизии. Методы решения и какой метод используется в го?

// Разрешение коллизий в хеш-функциях
// http://genius.pstu.ru/file.php/1/pupils_works_2017/MuhinaAlisa.pdf

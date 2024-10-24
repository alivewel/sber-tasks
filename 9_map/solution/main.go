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


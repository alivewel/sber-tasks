package main

import (
	"runtime"
)

// что выведет программа?

func main() {
	ch := make(chan string)
	close(ch)
	go func() {
		text := <-ch
		println("Hello, ", text)
	}()
	runtime.GC()
}

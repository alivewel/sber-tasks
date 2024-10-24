package main

import (
	"fmt"
	"sync"
)

const N = 5

func main() {
	in := make([]chan int, 0)
	for i := 0; i < N; i++ {
		in = append(in, make(chan int))
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		cc := make([]<-chan int, 0)
		for _, c := range in {
			cc = append(cc, c)
		}
		for v := range merge(cc...) {
			fmt.Println(v)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 20; i++ {
			in[i%N] <- i
		}
		for _, c := range in {
			close(c)
		}
		wg.Done()
	}()

	wg.Wait()
}

func merge(in ...<-chan int) <-chan int {
	// Let's do a better solution
	return nil
}

package pipeline

import "fmt"

// Assignment 5: Pipeline Pattern Problem: Implement a pipeline pattern in Go where
// one goroutine generates numbers, another squares them, and a
// third prints the squared numbers.
// print(square(generate(n)))

func Generate(n int) chan int {
	channel := make(chan int)
	go func() {
		for i := 1; i < n; i++ {
			channel <- i
		}
		close(channel)
	}()
	return channel
}

func Square(channel chan int) chan int {
	sq_channel := make(chan int)
	go func() {
		for x := range channel {
			sq_channel <- (x * x)
		}
		close(sq_channel)
	}()
	return sq_channel
}

func Print(channel chan int) {
	for num := range channel {
		fmt.Println("x:", num)
	}
}

package main

import "fmt"

func main() {
	input := 5
	fmt.Println("Input of, ", input, " is ", fib(input))

	len := 45
	jobs := make(chan int, len)
	results := make(chan int, len)

	go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)

	for i := 0; i < len; i++ {
		jobs <- i
	}
	close(jobs)
	for j := 0; j < len; j++ {
		fmt.Println(<-results)
	}

}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}

}

func fib(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2) + 1

}

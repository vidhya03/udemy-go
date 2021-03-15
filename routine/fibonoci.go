package main

import "fmt"

func main() {
	fmt.Println(fib(5))
}

func worker(jobs <-chan int, results <-chan int) {

}

func fib(n int) int {
	if n == 1 {
		return n
	} else if n <= 0 {
		return 0
	}
	return fib(n-1) + fib(n-2)

}

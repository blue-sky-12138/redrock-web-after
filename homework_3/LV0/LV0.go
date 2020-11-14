package main

import (
	"fmt"
)

var (
	value = make(chan int)
)

func factorial(n int,value chan int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	value <- res
}

func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i,value)
		fmt.Printf("myres[%d] = %d\n", i, <-value)
	}
	close(value)
}
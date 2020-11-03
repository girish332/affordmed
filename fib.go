package main

import (
	"fmt"
)

var cache = make(map[int]int)

func fib(n int, c chan int) {
	cache[0] = 0
	cache[1] = 1

	if n >= 0 {

		val, status := cache[n]

		if status == true {
			c <- val
			// return val
		} else {

			for i := len(cache); i <= n; i++ {

				cache[i] = cache[i-1] + cache[i-2]
			}
			c <- cache[n]
			// return cache[n]
		}

	} else {
		c <- -1
		// return -1
	}

}

func main() {

	c := make(chan int)
	go fib(10, c)
	ans := <-c
	// ans := fib(10)

	if ans >= 0 {
		fmt.Println(ans)

	} else {
		fmt.Println("Cannot calculate fibonacci number for negative index")
	}
}

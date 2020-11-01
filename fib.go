package main

import (
	"fmt"
)

var cache = make(map[int]int)

func fib(n int) int {
	cache[0] = 0
	cache[1] = 1

	if n >= 0 {

		val, status := cache[n]

		if status == true {

			return val
		} else {

			for i := 2; i <= n; i++ {

				cache[i] = cache[i-1] + cache[i-2]
			}

			return cache[n]
		}

	} else {
		return -1
	}

}

func main() {

	ans := fib(100)

	if ans >= 0 {
		fmt.Println(ans)

	} else {
		fmt.Println("Cannot calculate fibonacci number for negative index")
	}
}

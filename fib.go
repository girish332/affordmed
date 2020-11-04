package main

import (
	"fmt"
	"sync"
)

type cache1 struct {
	mu sync.Mutex
	v  map[int]int
}

var c = cache1{v: make(map[int]int)}

func fib1(n int) int {

	c.v[0] = 0
	c.v[1] = 1
	ans := -1
	if n >= 0 {
		c.mu.Lock()

		val, status := c.v[n]

		if status == true {

			ans = val
			c.mu.Unlock()
		} else {

			for i := len(c.v); i <= n; i++ {
				c.v[i] = c.v[i-1] + c.v[i-2]
			}
			ans = c.v[n]

			defer c.mu.Unlock()
		}
	}

	return ans

}

func main() {

	ans := fib1(10)
	if ans == -1 {
		fmt.Println("Cannot find fibo number for negative index")
	} else {

		fmt.Println(ans)
	}

}

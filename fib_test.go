package main

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {

	var tests = []struct {
		a    int
		want int
	}{
		{0, 0},
		{10, 55},
		{30, 832040},
		{70, 190392490709135},
		{-1, -1},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {

			ans := fib(tt.a)
			if ans != tt.want {
				t.Errorf("got %d,want %d", ans, tt.want)
			}
		})

	}

}

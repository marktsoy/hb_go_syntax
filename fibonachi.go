package main

import (
	"fmt"
	"time"
)

func trackTime(fn func()) {
	from := time.Now()
	fn()
	to := time.Now()
	elapsed := to.Sub(from)
	fmt.Printf("Elapsed: %v\n", elapsed.Microseconds())
}

func memoizedFib() func(uint) uint64 {
	var mem = []uint64{0, 1}

	var fib func(uint) uint64
	fib = func(i uint) uint64 {
		if i >= uint(len(mem)) {
			val := fib(i-1) + fib(i-2)
			mem = append(mem, val)
		}
		return mem[i]
	}
	return fib
}

func basic_fibonachi(i uint) uint64 {
	if i <= 1 {
		return uint64(i)
	}
	val := basic_fibonachi(i-1) + basic_fibonachi(i-2)

	return val
}

func main() {
	fibonachi := memoizedFib()
	fmt.Println("Basic fibonachi")
	trackTime(func() {
		var i uint = 0
		for ; i < 40; i++ {
			fmt.Println(basic_fibonachi(i))
		}
	})

	fmt.Println("Memoized fibonachi")
	trackTime(func() {
		var i uint = 0
		for ; i < 40; i++ {
			fmt.Println(fibonachi(i))
		}
	})
}

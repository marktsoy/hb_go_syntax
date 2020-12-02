package main

import "fmt"

func oddEven(elems []int) (int, int) {
	var odd, even int
	for _, n := range elems {
		if n%2 == 0 {
			even += n
		} else {
			odd += n
		}
	}
	return odd, even
}

/**
* Tring function as values
**/
func oddEventFunc() func(n int) (odd int, even int) {
	odds, evens := 0, 0
	return func(n int) (int, int) {
		if n%2 == 0 {
			evens += n
		} else {
			odds += n
		}
		return odds, evens
	}
}

func main() {
	fmt.Println("Basic odd even")
	elements := []int{100, 101, 100, 200, 201}
	o, e := oddEven(elements)
	fmt.Printf("Odd Sum:\t%v \nEven Sum:\t %v\n", o, e)

	fmt.Println("Func -> Func")
	push := oddEventFunc()
	for _, el := range elements {
		push(el)
	}
	odd, even := push(0)
	fmt.Printf("Odd Sum:\t%v \nEven Sum:\t %v\n", odd, even)

}

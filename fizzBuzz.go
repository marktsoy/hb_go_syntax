/**
* Divisible by 3 Buzz
* Other Fizz
**/
package main

import "fmt"

func fizzBuzz(max uint) {
	var i uint
	for i < max {
		i++

		switch {

		case i%15 == 0:
			{
				fmt.Printf("%v, ", "FizzBuzz")
			}
		case i%3 == 0:
			{
				fmt.Printf("%v, ", "Fizz")
			}

		case i%5 == 0:
			{
				fmt.Printf("%v, ", "Buzz")
			}
		default:
			{
				fmt.Printf("%v, ", i)
			}
		}
	}
}

func main() {
	fizzBuzz(100)
}

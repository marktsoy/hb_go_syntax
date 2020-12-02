package main

import "fmt"

func countElements(sl []int) map[int]int {
	var counter = make(map[int]int)
	for _, el := range sl {
		counter[el]++
	}

	return counter
}

func main() {
	elems := []int{1, 2, 3, 4, 4, 5, 1, 22, 3, 4, 55, 6, 12, 34, 56, 7, 3, 4, 5, 6, 7}
	filtered := []int{}
	uniques := countElements(elems)

	for number, count := range uniques {
		filtered = append(filtered, number)
		if count > 1 {
			fmt.Printf("Element: %v  \t occured %v times\n", number, count)
		}
	}
	fmt.Println(elems)
	fmt.Println(filtered)
}

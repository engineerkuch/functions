// A binary search function.
package main

import "fmt"

func main() {

	// testing our function with an array of integers.
	list := []int{1, 2, 3, 4, 5, 6, 10, 77, 500, 7, 15}
	// capacity := cap(list) - 1
	item := 77
	searchList := binarySearch(list, item)

	if searchList != -1 {
		fmt.Println("item found")
	} else {
		fmt.Println("item not found")
	}

}

// Note: Only simple integers
func binarySearch(list []int, item int) int {

	left := 0
	right := cap(list) - 1

	for left <= right {
		midpoint := (left + right) / 2
		if item == list[midpoint] {
			return midpoint
		} else if item > list[midpoint] {
			left = midpoint + 1
		} else {
			right = midpoint - 1
		}
	}

	return -1		// The binary function returns -1 if no item is found. 
}

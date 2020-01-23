// A binary search function.
package main

import "fmt"

func main() {

	// testing our function with an array of integers.
	list := []int{1, 2, 3, 4, 5, 6, 10, 77, 500, 7, 15}
	capacity := cap(list) - 1
	item := 77
	searchList := binarySearch(list, capacity, item)

	if searchList != -1 {
		fmt.Println("item found")
	} else {
		fmt.Println("item not found")
	}

}

// integer version
func binarySearch(list []int, size int, item int) int {

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

	return -1
}

// string version.
// func binarySearch(list []string, size int, item string) string {

// 	left := 0
// 	right := cap(list) - 1

// 	for left <= right {
// 		midpoint := (left + right) / 2
// 		if item == list[midpoint] {
// 			return midpoint
// 		} else if item > list[midpoint] {
// 			left = midpoint + 1
// 		} else {
// 			right = midpoint - 1
// 		}
// 	}

// 	return -1
// }

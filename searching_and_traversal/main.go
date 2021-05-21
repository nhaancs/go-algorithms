package main

import "fmt"

func main() {
	// fmt.Println(LinearSearch([]int{}, 5))
	// fmt.Println(LinearSearch([]int{1}, 5))
	// fmt.Println(LinearSearch([]int{6}, 5))
	// fmt.Println(LinearSearch([]int{5}, 5))
	// fmt.Println(LinearSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
	fmt.Println(BinarySearch([]int{}, 5))
	fmt.Println(BinarySearch([]int{1}, 5))
	fmt.Println(BinarySearch([]int{6}, 5))
	fmt.Println(BinarySearch([]int{5}, 5))
	fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
}

package main

import (
	"fmt"
	"go_data_structures_algorithms/searching_and_traversal"
)

func main() {
	// fmt.Println(searching_and_traversal.LinearSearch([]int{}, 5))
	// fmt.Println(searching_and_traversal.LinearSearch([]int{1}, 5))
	// fmt.Println(searching_and_traversal.LinearSearch([]int{6}, 5))
	// fmt.Println(searching_and_traversal.LinearSearch([]int{5}, 5))
	// fmt.Println(searching_and_traversal.LinearSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
	fmt.Println(searching_and_traversal.BinarySearch([]int{}, 5))
	fmt.Println(searching_and_traversal.BinarySearch([]int{1}, 5))
	fmt.Println(searching_and_traversal.BinarySearch([]int{6}, 5))
	fmt.Println(searching_and_traversal.BinarySearch([]int{5}, 5))
	fmt.Println(searching_and_traversal.BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
}

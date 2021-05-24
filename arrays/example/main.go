package main

import (
	"fmt"
	"go_data_structures_algorithms/arrays"
)

func main() {
	// arr := arrays.NewMyStringArray(5)
	// fmt.Println("initial array: ", arr)
	// arr.Insert(3, "World")
	// arr.Insert(0, "Hello")
	// arr.Insert(4, "Demo")
	// arr.Delete(4)
	// fmt.Println("array search: ", arr.Search("Demo"))
	// fmt.Println(arr.Access(0), arr)

	// fmt.Println(Reverse("hello world!"))
	fmt.Println(arrays.MergeSortedArrays([]int{4, 6, 8}, []int{3, 5, 7, 9}))
}

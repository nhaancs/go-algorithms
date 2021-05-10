package main

import (
	"arrays/my_array"
	"fmt"
)

func main() {
	arr := my_array.NewMyStringArray(5)
	fmt.Println("initial array: ", arr)
	arr.Insert(3, "World")
	arr.Insert(0, "Hello")
	arr.Insert(4, "Demo")
	arr.Delete(4)
	fmt.Println("array search: ", arr.Search("Demo"))
	fmt.Println(arr.Access(0), arr)
}
package main

import "fmt"

func main() {
	// stll := NewStackLinkedList()
	// stll.Push(1)
	// stll.Push(2)
	// stll.Push(3)
	// stll.Push(4)
	// stll.Print()
	// stll.Push(5)
	// stll.Print()
	// fmt.Println(stll.Peek())
	// fmt.Println(stll.Pop())
	// fmt.Println(stll.Pop())
	// fmt.Println(stll.Pop())
	// stll.Print()
	
	starr := NewStackArray()
	starr.Push(1)
	starr.Push(2)
	starr.Push(3)
	starr.Push(4)
	starr.Print()
	starr.Push(5)
	starr.Print()
	fmt.Println(starr.Peek())
	fmt.Println(starr.Pop())
	fmt.Println(starr.Pop())
	fmt.Println(starr.Pop())
	starr.Print()
}

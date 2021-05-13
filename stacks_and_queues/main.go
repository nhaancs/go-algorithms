package main

import "fmt"

func main() {
	stll := NewStackLinkedList()
	stll.Push(1)
	stll.Push(2)
	stll.Push(3)
	stll.Push(4)
	stll.Print()
	stll.Push(5)
	stll.Print()
	fmt.Println(stll.Peek())
	fmt.Println(stll.Pop())
	fmt.Println(stll.Pop())
	fmt.Println(stll.Pop())
	stll.Print()
}

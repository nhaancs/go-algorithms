package main

import (
	"github.com/nhaancs/go-algorithms/archived/linked_lists"
)

func main() {
	// linkedList := linked_lists.NewLinkedList()
	// linkedList.Append(10)
	// linkedList.Append(11)
	// linkedList.Append(12)
	// linkedList.Append(13)
	// linkedList.Print()
	// linkedList.Prepend(5)
	// linkedList.Print()
	// linkedList.Insert(2, 9)
	// linkedList.Insert(0, 8)
	// linkedList.Insert(20, 88)
	// linkedList.Print()
	// linkedList.Remove(0)
	// linkedList.Print()

	// linkedList2 := linked_lists.NewLinkedList2()
	// linkedList2.Append(1)
	// linkedList2.Append(2)
	// linkedList2.Append(3)
	// linkedList2.Prepend(-1)
	// linkedList2.Prepend(-2)
	// linkedList2.Prepend(-3)
	// linkedList2.Insert(3, 0)
	// linkedList2.Remove(3)
	// linkedList2.Print()

	linkedList := linked_lists.NewLinkedList()
	linkedList.Append(10)
	linkedList.Append(11)
	linkedList.Append(12)
	linkedList.Append(13)
	linkedList.Print()
	linkedList.Reverse()
	linkedList.Print()
}

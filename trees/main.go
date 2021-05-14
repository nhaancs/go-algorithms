package main

import "fmt"

func main() {
	bst := NewBinarySearchTree()
	bst.Insert(10)
	bst.Insert(8)
	bst.Insert(9)
	bst.Insert(11)
	// bst.Remove(9)
	// fmt.Println(bst.Insert(8))
	fmt.Println(bst.Lookup(9))
}

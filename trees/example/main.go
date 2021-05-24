package main

import (
	"fmt"
	"go_data_structures_algorithms/trees"
)

func main() {
	// bst := trees.NewBinarySearchTree()
	// bst.Insert(10)
	// bst.Insert(8)
	// bst.Insert(9)
	// bst.Insert(11)
	// // bst.Remove(9)
	// // fmt.Println(bst.Insert(8))
	// fmt.Println(bst.Lookup(9))

	var rbst trees.RBinarySearchTree
	rbst.Insert(10, "10")
	rbst.Insert(8, "8")
	rbst.Insert(9, "9")
	rbst.Insert(11, "11")
	rbst.Insert(8, "8")
	rbst.Insert(12, "12")
	rbst.Insert(20, "20")
	// fmt.Println(rbst.Search(9))
	rbst.String()

	fmt.Println("Pre-order traverse")
	rbst.PreOrderTraverse(func(v interface{}) {
		fmt.Print(v, " ")
	})
	fmt.Println()

	fmt.Println("In-order traverse")
	rbst.InOrderTraverse(func(v interface{}) {
		fmt.Print(v, " ")
	})
	fmt.Println()

	fmt.Println("Post-order traverse")
	rbst.PostOrderTraverse(func(v interface{}) {
		fmt.Print(v, " ")
	})
	fmt.Println()
}

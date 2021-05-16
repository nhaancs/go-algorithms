package main

func main() {
	// bst := NewBinarySearchTree()
	// bst.Insert(10)
	// bst.Insert(8)
	// bst.Insert(9)
	// bst.Insert(11)
	// // bst.Remove(9)
	// // fmt.Println(bst.Insert(8))
	// fmt.Println(bst.Lookup(9))
	
	var rbst ItemBinarySearchTree
	rbst.Insert(10, "10")
	rbst.Insert(8, "8")
	rbst.Insert(9, "9")
	rbst.Insert(11, "11")
	rbst.Insert(8, "8")
	// fmt.Println(rbst.Search(9))
	rbst.String()
}

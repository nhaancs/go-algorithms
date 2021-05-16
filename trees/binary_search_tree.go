package main

import "errors"

type bstNode struct {
	Value int
	Left  *bstNode
	Right *bstNode
}

func NewBTSNode(value int) *bstNode {
	return &bstNode{Value: value}
}

type binarySearchTree struct {
	Root *bstNode
}

func NewBinarySearchTree() *binarySearchTree {
	return &binarySearchTree{}
}

func (bst *binarySearchTree) Insert(value int) error {
	newNode := NewBTSNode(value)
	if bst.Root == nil {
		bst.Root = newNode
		return nil
	}

	currentNode := bst.Root
	for {
		if value < currentNode.Value { // left
			if currentNode.Left == nil {
				currentNode.Left = newNode
				return nil
			}

			currentNode = currentNode.Left
			continue
		}

		if value > currentNode.Value { // Right
			if currentNode.Right == nil {
				currentNode.Right = newNode
				return nil
			}

			currentNode = currentNode.Right
			continue
		}

		if value == currentNode.Value {
			return errors.New("duplicated value")
		}
	}
}

func (bst *binarySearchTree) Lookup(value int) *bstNode {
	currentNode := bst.Root
	for currentNode != nil {
		if value < currentNode.Value { // left
			currentNode = currentNode.Left
			continue
		}

		if value > currentNode.Value { // Right
			currentNode = currentNode.Right
			continue
		}

		if value == currentNode.Value {
			return currentNode
		}
	}

	return nil
}

// todo: recheck logic??
func (bst *binarySearchTree) Remove(value int) {
	var parentNode *bstNode
	currentNode := bst.Root
	for currentNode != nil {
		if value < currentNode.Value {
			parentNode = currentNode
			currentNode = currentNode.Left
			continue
		}

		if value > currentNode.Value {
			parentNode = currentNode
			currentNode = currentNode.Right
			continue
		}

		// value == currentNode.Value
		// we have a match, get to work!

		// option 1: no right child // only left
		if currentNode.Right == nil {
			if parentNode == nil { // delete root node, root node does not have any right child
				bst.Root = currentNode.Left // make the left child of the current root become the next root
			} else {
				// if current < parent value, make current left child a left child of parent
				if currentNode.Value < parentNode.Value {
					// parentNode.Left is the currentNode in this case!!!
					// currentNode is less than parent node => currentNode.Left is also less than parentNode
					parentNode.Left = currentNode.Left
					// if current > parent value, make current left child a right child of parent
				} else if currentNode.Value > parentNode.Value {
					// parentNode.Right is the currentNode in this case!!!
					// currentNode is greater than parent node => currentNode.Left is also greater than parentNode
					parentNode.Right = currentNode.Left
				}
			}
			// option 2: right child which does not have a left child // only right
		} else if currentNode.Right.Left == nil {
			// now assign current left child to a left child of right child
			currentNode.Right.Left = currentNode.Left
			if parentNode == nil {
				bst.Root = currentNode.Right
			} else {
				// parentNode.Left is the currentNode in this case!!!
				// if current < parent, make right child of the left the parent
				if currentNode.Value < parentNode.Value {
					parentNode.Left = currentNode.Right
					// parentNode.Right is the currentNode in this case!!!
					// if current > parent, make right child a right child of the parent
				} else if currentNode.Value > parentNode.Value {
					parentNode.Right = currentNode.Right
				}
			}
			// option 3: right child that has a left child // have both left and right
		} else {
			// find the right child's left most child
			leftMost := currentNode.Right.Left
			leftMostParent := currentNode.Right
			for leftMost.Left != nil {
				leftMostParent = leftMost
				leftMost = leftMost.Left
			}

			// parent's left subtree is now leftmost's right subtree
			leftMostParent.Left = leftMost.Right
			leftMost.Left = currentNode.Left
			leftMost.Right = currentNode.Right

			if parentNode == nil {
				bst.Root = leftMost
			} else {
				if currentNode.Value < parentNode.Value {
					parentNode.Left = leftMost
				} else if currentNode.Value > parentNode.Value {
					parentNode.Right = leftMost
				}
			}
		}

		return
	}
}

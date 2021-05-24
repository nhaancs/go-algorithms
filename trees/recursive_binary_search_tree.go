package trees

// https://flaviocopes.com/golang-data-structure-binary-search-tree

import (
	"errors"
	"fmt"
	"sync"
)

// Node a single node that composes the tree
type Node struct {
    key   int
    value interface{}
    left  *Node //left
    right *Node //right
}

// RBinarySearchTree the binary search tree of Items
type RBinarySearchTree struct {
    root *Node
    lock sync.RWMutex
}

// Insert inserts the interface{} t in the tree
func (bst *RBinarySearchTree) Insert(key int, value interface{}) {
    bst.lock.Lock()
    defer bst.lock.Unlock()
    n := &Node{key, value, nil, nil}
    if bst.root == nil {
        bst.root = n
    } else {
        insertNode(bst.root, n)
    }
}

// internal function to find the correct place for a node in a tree
func insertNode(node, newNode *Node) error {
    if newNode.key < node.key {
        if node.left == nil {
            node.left = newNode
        } else {
            insertNode(node.left, newNode)
        }
        return nil
    } 

    if newNode.key > node.key {
        if node.right == nil {
            node.right = newNode
        } else {
            insertNode(node.right, newNode)
        }
        return nil
    }
    
    return errors.New("duplicated key")
}

// InOrderTraverse visits all nodes with in-order traversing
func (bst *RBinarySearchTree) InOrderTraverse(f func(interface{})) {
    bst.lock.RLock()
    defer bst.lock.RUnlock()
    inOrderTraverse(bst.root, f)
}

// internal recursive function to traverse in order
func inOrderTraverse(n *Node, f func(interface{})) {
    if n != nil {
        inOrderTraverse(n.left, f)
        f(n.value)
        inOrderTraverse(n.right, f)
    }
}

// PreOrderTraverse visits all nodes with pre-order traversing
func (bst *RBinarySearchTree) PreOrderTraverse(f func(interface{})) {
    bst.lock.Lock()
    defer bst.lock.Unlock()
    preOrderTraverse(bst.root, f)
}

// internal recursive function to traverse pre order
func preOrderTraverse(n *Node, f func(interface{})) {
    if n != nil {
        f(n.value)
        preOrderTraverse(n.left, f)
        preOrderTraverse(n.right, f)
    }
}

// PostOrderTraverse visits all nodes with post-order traversing
func (bst *RBinarySearchTree) PostOrderTraverse(f func(interface{})) {
    bst.lock.Lock()
    defer bst.lock.Unlock()
    postOrderTraverse(bst.root, f)
}

// internal recursive function to traverse post order
func postOrderTraverse(n *Node, f func(interface{})) {
    if n != nil {
        postOrderTraverse(n.left, f)
        postOrderTraverse(n.right, f)
        f(n.value)
    }
}

// Min returns the interface{} with min value stored in the tree
func (bst *RBinarySearchTree) Min() *interface{} {
    bst.lock.RLock()
    defer bst.lock.RUnlock()
    n := bst.root
    if n == nil {
        return nil
    }
    for {
        if n.left == nil {
            return &n.value
        }
        n = n.left
    }
}

// Max returns the interface{} with max value stored in the tree
func (bst *RBinarySearchTree) Max() *interface{} {
    bst.lock.RLock()
    defer bst.lock.RUnlock()
    n := bst.root
    if n == nil {
        return nil
    }
    for {
        if n.right == nil {
            return &n.value
        }
        n = n.right
    }
}

// Search returns true if the interface{} t exists in the tree
func (bst *RBinarySearchTree) Search(key int) bool {
    bst.lock.RLock()
    defer bst.lock.RUnlock()
    return search(bst.root, key)
}

// internal recursive function to search an item in the tree
func search(n *Node, key int) bool {
    if n == nil {
        return false
    }
    if key < n.key {
        return search(n.left, key)
    }
    if key > n.key {
        return search(n.right, key)
    }
    return true
}

// Remove removes the interface{} with key `key` from the tree
func (bst *RBinarySearchTree) Remove(key int) {
    bst.lock.Lock()
    defer bst.lock.Unlock()
    remove(bst.root, key)
}

// internal recursive function to remove an item
func remove(node *Node, key int) *Node {
    if node == nil {
        return nil
    }
    if key < node.key {
        node.left = remove(node.left, key)
        return node
    }
    if key > node.key {
        node.right = remove(node.right, key)
        return node
    }
    // key == node.key

    // deleted node does not have right and left child
    if node.left == nil && node.right == nil {
        node = nil
        return nil
    }

    // deleted node does not have left child
    if node.left == nil {
        node = node.right
        return node
    }

    // deleted node does not have right child
    if node.right == nil {
        node = node.left
        return node
    }

    // deleted node does not have both left and right child
    leftmostrightside := node.right
    for {
        //find smallest value on the right side
        if leftmostrightside != nil && leftmostrightside.left != nil {
            leftmostrightside = leftmostrightside.left
        } else {
            break
        }
    }
    node.key, node.value = leftmostrightside.key, leftmostrightside.value
    // todo: recheck logic??
    node.right = remove(node.right, node.key)
    return node
}

// String prints a visual representation of the tree
func (bst *RBinarySearchTree) String() {
    bst.lock.Lock()
    defer bst.lock.Unlock()
    fmt.Println("------------------------------------------------")
    stringify(bst.root, 0)
    fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *Node, level int) {
    if n != nil {
        format := ""
        for i := 0; i < level; i++ {
            format += "       "
        }
        format += "---[ "
        level++
        stringify(n.left, level)
        fmt.Printf(format+"%d\n", n.key)
        stringify(n.right, level)
    }
}

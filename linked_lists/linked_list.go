package main

import "fmt"

type node struct {
	Value interface{}
	Next *node
}

func newNode(value interface{}) *node {
	return &node{Value: value}
}

type linkedList struct {
	Head *node
	Tail *node
	Length uint
}

func NewLinkedList() *linkedList {
	return &linkedList{}
}

func (ll *linkedList) Append(value interface{}) {
	node := newNode(value)
	if ll.Length == 0 {
		ll.Head = node
	} else {
		ll.Tail.Next = node
	}

	ll.Tail = node
	ll.Length++
}

func (ll *linkedList) Prepend(value interface{}) {
	node := newNode(value)
	if ll.Length == 0 {
		ll.Tail = node
	} else {
		node.Next = ll.Head
	}

	ll.Head = node
	ll.Length++
}

func (ll *linkedList) Insert(index uint, value interface{}) {
	if (index >= ll.Length) {
		ll.Append(value)
		return 
	}

	if (index == 0) {
		ll.Prepend(value)
		return
	}

	rightBeforeInsertedNode := ll.Get(index-1)
	nodeAtInsertedIndex := rightBeforeInsertedNode.Next

	insertedNode := newNode(value)
	rightBeforeInsertedNode.Next = insertedNode
	insertedNode.Next = nodeAtInsertedIndex

	ll.Length++
}

func (ll *linkedList) Get(index uint) *node {
	if (index > ll.Length-1) {
		return nil
	}

	runner := ll.Head
	var i uint
	for ; i < index; i++ {
		runner = runner.Next
	}

	return runner
}

func (ll *linkedList) Remove(index uint) {
	if index > ll.Length-1 {
		return
	}

	if index == 0 {
		ll.Head = ll.Head.Next
		ll.Length--
		return
	}
	
	if index == ll.Length-1 {
		itemNextToLast := ll.Get(ll.Length-2)
		itemNextToLast.Next = nil
		ll.Tail = itemNextToLast
		ll.Length-- 
		return
	}
	
	rightBeforeRemovedNode := ll.Get(index-1)
	removedNode := rightBeforeRemovedNode.Next
	rightBeforeRemovedNode.Next = removedNode.Next

	ll.Length--
}

func (ll *linkedList) Print() {
	arr := []interface{}{}
	for runner := ll.Head; runner != nil; runner = runner.Next {
		arr = append(arr, runner.Value)
	}

	fmt.Println(arr)
}


package linked_lists

import "fmt"

type node2 struct {
	Value interface{}
	Next *node2
	Prev *node2
}

func newNode2(value interface{}) *node2 {
	return &node2{Value: value}
}

type linkedList2 struct {
	Head *node2
	Tail *node2
	Length uint
}

func NewLinkedList2() *linkedList2 {
	return &linkedList2{}
}

func (ll *linkedList2) Append(value interface{}) {
	node := newNode2(value)
	if ll.Length == 0 {
		ll.Head = node
	} else {
		node.Prev = ll.Tail
		ll.Tail.Next = node
	}

	ll.Tail = node
	ll.Length++
}

func (ll *linkedList2) Prepend(value interface{}) {
	node := newNode2(value)
	if ll.Length == 0 {
		ll.Tail = node
	} else {
		ll.Head.Prev = node
		node.Next = ll.Head
	}

	ll.Head = node
	ll.Length++
}

func (ll *linkedList2) Insert(index uint, value interface{}) {
	if (index >= ll.Length) {
		ll.Append(value)
		return 
	}

	if (index == 0) {
		ll.Prepend(value)
		return
	}

	
	nodeAtIndex := ll.Get(index)
	beforeNode := nodeAtIndex.Prev
	insertedNode := newNode2(value)

	beforeNode.Next = insertedNode
	insertedNode.Prev = beforeNode

	insertedNode.Next = nodeAtIndex
	nodeAtIndex.Prev = insertedNode

	ll.Length++
}

func (ll *linkedList2) Get(index uint) *node2 {
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

func (ll *linkedList2) Remove(index uint) {
	if index > ll.Length-1 {
		return
	}

	if index == 0 {
		ll.Head = ll.Head.Next
		ll.Length--
		return
	}
	
	if index == ll.Length-1 {
		ll.Tail = ll.Tail.Prev
		ll.Tail.Next = nil
		ll.Length-- 
		return
	}
	
	removedNode := ll.Get(index)
	removedNode.Prev.Next = removedNode.Next
	removedNode.Next.Prev = removedNode.Next
	ll.Length--
}

func (ll *linkedList2) Print() {
	arr := []interface{}{}
	
	for runner := ll.Head; runner != nil; runner = runner.Next {
		arr = append(arr, *runner)
	}

	fmt.Println(arr)
}


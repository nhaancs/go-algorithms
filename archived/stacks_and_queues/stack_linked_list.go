package stacks_and_queues

import "fmt"

type stackNode struct {
	value interface{}
	next  *stackNode
}

func newNode(value interface{}) *stackNode {
	return &stackNode{value: value}
}

type stackLinkedList struct {
	top    *stackNode // is the head of the linked list
	bottom *stackNode // is the tail of the linked list
	length int
}

func NewStackLinkedList() *stackLinkedList {
	return &stackLinkedList{}
}

func (stll *stackLinkedList) Peek() *stackNode {
	return stll.top
}

func (stll *stackLinkedList) Push(value interface{}) {
	newTop := newNode(value)
	if stll.length == 0 {
		stll.top = newTop
		stll.bottom = newTop
		stll.length++
		return
	}

	lastTop := stll.top
	newTop.next = lastTop
	stll.top = newTop
	stll.length++
}

func (stll *stackLinkedList) Pop() *stackNode {
	if stll.length == 0 {
		return nil
	}

	lastTop := stll.top
	stll.top = lastTop.next
	if stll.length == 1 {
		stll.bottom = nil
	}
	stll.length--
	return lastTop
}

func (stll *stackLinkedList) IsEmpty() bool {
	return stll.length == 0
}

func (stll *stackLinkedList) Print() {
	result := []interface{}{}
	for runner := stll.top; runner != nil; runner = runner.next {
		result = append(result, runner.value)
	}

	fmt.Println(result)
}

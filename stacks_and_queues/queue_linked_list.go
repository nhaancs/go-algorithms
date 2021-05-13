package main

import "fmt"

type queueNode struct {
	value interface{}
	next  *queueNode
}

func newQueueNode(value interface{}) *queueNode {
	return &queueNode{value: value}
}

type queueLinkedList struct {
	first  *queueNode
	last   *queueNode
	length int
}

func NewQueueLinkedList() *queueLinkedList {
	return &queueLinkedList{}
}

func (qll *queueLinkedList) Peek() *queueNode {
	return qll.first
}

func (qll *queueLinkedList) Enqueue(value interface{}) {
	newLast := newQueueNode(value)
	if qll.length == 0 {
		qll.first = newLast
	} else {
		qll.last.next = newLast
	}
	qll.last = newLast
	qll.length++
}

func (qll *queueLinkedList) Dequeue() *queueNode {
	if qll.length == 0 {
		return nil
	}

	lastFirst := qll.first
	if qll.length == 1 {
		qll.last = nil
	}
	qll.first = qll.first.next
	qll.length--

	return lastFirst
}

func (qll *queueLinkedList) IsEmpty() bool {
	return qll.length == 0
}

func (qll *queueLinkedList) Print() {
	result := []interface{}{}
	for runner := qll.first; runner != nil; runner = runner.next {
		result = append(result, runner.value)
	}

	fmt.Println(result)
}

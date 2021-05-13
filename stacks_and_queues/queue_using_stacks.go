package main

/*
Queue
- enqueue to last
- dequeue the first element
- peek fist element

Stack
- push to last
- pop the last element
- peek last element
*/

// https://leetcode.com/problems/implement-queue-using-stacks/solution/

type queueStack struct {
	s1 *stackLinkedList // main queue, s1.Pop() will remove and return the first element of the queue
	s2 *stackLinkedList // temp stack
}

func NewQueueStack() *queueStack {
	return &queueStack{s1: NewStackLinkedList(), s2: NewStackLinkedList()}
}

func (qs *queueStack) Peek() interface{} {
	return qs.s1.Peek().value
}

func (qs *queueStack) Enqueue(value interface{}) {
	// move all elements from s1 to s2
	for !qs.s1.IsEmpty() {
		qs.s2.Push(qs.s1.Pop().value)
	}

	// push new element to s1
	qs.s1.Push(value)

	// move all elements from s2 to s1
	for !qs.s2.IsEmpty() {
		qs.s1.Push(qs.s2.Pop().value)
	}
}

func (qs *queueStack) Dequeue() interface{} {
	// pop 1 element from s1
	return qs.s1.Pop().value
}

func (qs *queueStack) IsEmpty() bool {
	return qs.s1.IsEmpty()
}

func (qs *queueStack) Print() {
	qs.s1.Print()
}

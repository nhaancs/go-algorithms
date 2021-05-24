package stacks_and_queues

import "fmt"

type stackArray struct {
	data []interface{}
}

func NewStackArray() *stackArray {
	return &stackArray{}
}

func (starr *stackArray) Peek() interface{} {
	if len(starr.data) == 0 {
		return nil
	}
	return starr.data[len(starr.data)-1]
}

func (starr *stackArray) Push(value interface{}) {
	starr.data = append(starr.data, value)
}

func (starr *stackArray) Pop() interface{} {
	totalItems := len(starr.data)
	if totalItems == 0 {
		return nil
	}

	lastTop := starr.data[totalItems-1]

	if totalItems == 1 {
		starr.data = starr.data[:0]
	} else {
		starr.data = starr.data[:totalItems-1]
	}
	
	return lastTop
}

func (starr *stackArray) IsEmpty() bool {
	return len(starr.data) == 0
}

func (starr *stackArray) Print() {
	fmt.Println(starr.data...)
}

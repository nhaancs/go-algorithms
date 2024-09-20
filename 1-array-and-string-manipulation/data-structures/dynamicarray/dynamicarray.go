package dynamicarray

import (
	"fmt"
	"sync"
)

type stringDynamicArr struct {
	size     uint
	capacity uint
	data     map[uint]string
	mu       sync.Mutex
}

func New() *stringDynamicArr {
	return &stringDynamicArr{
		capacity: 1,
		data:     make(map[uint]string, 1),
	}
}

func (s *stringDynamicArr) Lookup(index uint) string {
	if index >= s.size {
		panic("index out of range")
	}
	return s.data[index]
}

func (s *stringDynamicArr) Append(value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.size >= s.capacity {
		s.resize()
	}
	s.data[s.size] = value
	s.size++
}

func (s *stringDynamicArr) Insert(index uint, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index >= s.size {
		panic("index out of range")
	}

	if s.size >= s.capacity {
		s.resize()
	}

	for i := s.size; i > index; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[index] = value
	s.size++
}

func (s *stringDynamicArr) Delete(index uint) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index >= s.size {
		panic("index out of range")
	}

	if index >= s.size {
		return
	}

	for i := index; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	s.data[s.size-1] = ""
	s.size--
}

func (s *stringDynamicArr) resize() {
	s.capacity *= 2
	newData := make(map[uint]string, s.capacity)
	for i := uint(0); i < s.size; i++ {
		newData[i] = s.data[i]
	}
	s.data = newData
}

func (s *stringDynamicArr) Size() uint {
	return s.size
}

func (s *stringDynamicArr) Capacity() uint {
	return s.capacity
}

func (s *stringDynamicArr) String() string {
	str := "[ "
	for i := uint(0); i < s.size; i++ {
		str += fmt.Sprintf("%d:%s ", i, s.data[i])
	}
	str += "]"
	return str
}

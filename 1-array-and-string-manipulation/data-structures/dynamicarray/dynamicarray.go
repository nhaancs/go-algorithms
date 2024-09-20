package dynamicarray

import (
	"fmt"
	"sync"
)

type stringDynamicArr struct {
	capacity uint
	data     map[uint]string
	mu       sync.Mutex
}

func New() *stringDynamicArr {
	return &stringDynamicArr{
		capacity: 1,
		data:     make(map[uint]string),
	}
}

func (s *stringDynamicArr) Lookup(index uint) string {
	if index >= s.Size() {
		panic("index out of range")
	}
	return s.data[index]
}

func (s *stringDynamicArr) Append(value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.Size() >= s.capacity {
		s.resize()
	}
	s.data[s.Size()] = value
}

func (s *stringDynamicArr) Insert(index uint, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index >= s.Size() {
		panic("index out of range")
	}

	if s.Size() >= s.capacity {
		s.resize()
	}

	for i := s.Size(); i > index; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[index] = value
}

func (s *stringDynamicArr) Delete(index uint) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index >= s.Size() {
		panic("index out of range")
	}

	for i := index; i < s.Size(); i++ {
		s.data[i] = s.data[i+1]
	}
	delete(s.data, s.Size()-1)
}

func (s *stringDynamicArr) resize() {
	s.capacity *= 2
	newData := make(map[uint]string, s.capacity)
	for i := uint(0); i < s.Size(); i++ {
		newData[i] = s.data[i]
	}
	s.data = newData
}

func (s *stringDynamicArr) Size() uint {
	return uint(len(s.data))
}

func (s *stringDynamicArr) Capacity() uint {
	return s.capacity
}

func (s *stringDynamicArr) String() string {
	str := "[ "
	for i := uint(0); i < s.Size(); i++ {
		str += fmt.Sprintf("%d:%s ", i, s.data[i])
	}
	str += "]"
	return str
}

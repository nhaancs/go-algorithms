package staticarray

import (
	"fmt"
	"sync"
)

type stringStaticArr struct {
	size     uint
	capacity uint
	data     map[uint]string
	mu       sync.Mutex
}

func New(capacity uint) *stringStaticArr {
	return &stringStaticArr{
		capacity: capacity,
		data:     make(map[uint]string, capacity),
	}
}

func (s *stringStaticArr) Lookup(index uint) string {
	if index >= s.capacity {
		panic("index out of range")
	}
	return s.data[index]
}

func (s *stringStaticArr) Append(value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.size >= s.capacity {
		panic("array is full")
	}
	s.data[s.size] = value
	s.size++
}

func (s *stringStaticArr) Insert(index uint, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index >= s.capacity {
		panic("index out of range")
	}

	if s.size >= s.capacity {
		panic("array is full")
	}

	for i := s.size; i > index; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[index] = value
	s.size++
}

func (s *stringStaticArr) Delete(index uint) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index >= s.capacity {
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

func (s *stringStaticArr) Size() uint {
	return s.size
}

func (s *stringStaticArr) Capacity() uint {
	return s.capacity
}

func (s *stringStaticArr) String() string {
	str := "[ "
	for i := uint(0); i < s.size; i++ {
		str += fmt.Sprintf("%d:%s ", i, s.data[i])
	}
	str += "]"
	return str
}

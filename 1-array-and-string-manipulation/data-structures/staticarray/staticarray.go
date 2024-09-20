package staticarray

import (
	"fmt"
	"sync"
)

type stringStaticArr struct {
	capacity uint
	data     map[uint]string
	mu       sync.Mutex
}

func New(capacity uint) *stringStaticArr {
	return &stringStaticArr{
		capacity: capacity,
		data:     make(map[uint]string),
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

	if s.size() >= s.capacity {
		panic("array is full")
	}
	s.data[s.size()] = value
}

func (s *stringStaticArr) Insert(index uint, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index >= s.capacity {
		panic("index out of range")
	}

	if s.size() >= s.capacity {
		panic("array is full")
	}

	for i := s.size(); i > index; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[index] = value
}

func (s *stringStaticArr) Delete(index uint) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index >= s.capacity {
		panic("index out of range")
	}

	if index >= s.size() {
		return
	}

	for i := index; i < s.size(); i++ {
		s.data[i] = s.data[i+1]
	}
	delete(s.data, s.size()-1)
}

func (s *stringStaticArr) size() uint {
	return uint(len(s.data))
}

func (s *stringStaticArr) Capacity() uint {
	return s.capacity
}

func (s *stringStaticArr) String() string {
	str := "[ "
	for i := uint(0); i < s.size(); i++ {
		str += fmt.Sprintf("%d:%s ", i, s.data[i])
	}
	str += "]"
	return str
}

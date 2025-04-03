package main

import "fmt"

type Stack8 struct {
	items []int
}

func (s *Stack8) Push(element int) {
	s.items = append(s.items, element)
}

func (s *Stack8) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

func (s *Stack8) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack8) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := Stack8{}
	stack.Push(10)
	stack.Push(20)
	fmt.Println(stack.Pop())  // 20, true
	fmt.Println(stack.Peek()) // 10, true
}

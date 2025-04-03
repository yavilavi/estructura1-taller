package main

import "fmt"

type Queue8 struct {
	items []string
}

func (q *Queue8) Enqueue(element string) {
	q.items = append(q.items, element)
}

func (q *Queue8) Dequeue() (string, bool) {
	if len(q.items) == 0 {
		return "", false
	}
	first := q.items[0]
	q.items = q.items[1:]
	return first, true
}

func (q *Queue8) Front() (string, bool) {
	if len(q.items) == 0 {
		return "", false
	}
	return q.items[0], true
}

func main() {
	Queue8 := Queue8{}
	Queue8.Enqueue("A")
	Queue8.Enqueue("B")
	fmt.Println(Queue8.Dequeue()) // "A", true
	fmt.Println(Queue8.Front())   // "B", true
}

package main

import (
	"fmt"
	"sort"
)

// Listas (Slices)
func listOperations() {
	array := []int{3, 1, 4, 1, 5}
	array = append(array, 9)                // Inserción
	array = append(array[:1], array[2:]...) // Borrado
	array[1] = 2                            // Actualización
	sort.Ints(array)                        // Ordenación
	fmt.Println("Lista ordenada:", array)
}

// Conjuntos (Maps como Sets)
func setOperations() {
	setData := map[int]bool{3: true, 1: true, 4: true, 5: true}
	setData[9] = true  // Inserción
	delete(setData, 1) // Borrado
	fmt.Println("Conjunto:", setData)
}

// Diccionarios (Maps)
func mapOperations() {
	mapData := map[string]int{"uno": 1, "dos": 2}
	mapData["tres"] = 3    // Inserción
	delete(mapData, "dos") // Borrado
	mapData["tres"] = 33   // Actualización
	fmt.Println("Mapa:", mapData)
}

// Pila (Stack) con Slice
type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

//s := Stack{}
//s.Push(10)
//s.Push(20)
//fmt.Println(s.Pop()) // 20
//fmt.Println(s.Pop()) // 10
//fmt.Println(s.Pop()) // -1 (stack is empty)

// Cola (Queue) con Slice
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

//q := Queue{}
//q.Enqueue(10)
//q.Enqueue(20)
//fmt.Println(q.Dequeue()) // 10
//fmt.Println(q.Dequeue()) // 20
//fmt.Println(q.Dequeue()) // -1 (queue is empty)

func main() {

}

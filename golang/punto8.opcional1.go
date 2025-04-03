package main

import "fmt"

type Stack8_1 struct {
	items []string
}

func (s *Stack8_1) Push(element string) {
	s.items = append(s.items, element)
}

func (s *Stack8_1) Pop() (string, bool) {
	if len(s.items) == 0 {
		return "", false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

type BrowserHistory struct {
	backStack    Stack8_1
	forwardStack Stack8_1
	current      string
}

func (b *BrowserHistory) Visit(page string) {
	if b.current != "" {
		b.backStack.Push(b.current)
	}
	b.current = page
	b.forwardStack = Stack8_1{}
	fmt.Println("Visitando:", page)
}

func (b *BrowserHistory) GoBack() {
	if val, ok := b.backStack.Pop(); ok {
		b.forwardStack.Push(b.current)
		b.current = val
		fmt.Println("Retrocediendo a:", b.current)
	} else {
		fmt.Println("No hay más páginas atrás")
	}
}

func (b *BrowserHistory) GoForward() {
	if val, ok := b.forwardStack.Pop(); ok {
		b.backStack.Push(b.current)
		b.current = val
		fmt.Println("Avanzando a:", b.current)
	} else {
		fmt.Println("No hay más páginas adelante")
	}
}

func main() {
	browser := BrowserHistory{}
	browser.Visit("google.com")
	browser.Visit("youtube.com")
	browser.GoBack()    // Retrocediendo a google.com
	browser.GoForward() // Avanzando a youtube.com
}

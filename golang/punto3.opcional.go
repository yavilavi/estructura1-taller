package main

import "fmt"

func procesarNumeros(texto1, texto2 string) int {
	contador := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(texto1 + texto2)
		} else if i%3 == 0 {
			fmt.Println(texto1)
		} else if i%5 == 0 {
			fmt.Println(texto2)
		} else {
			fmt.Println(i)
			contador++
		}
	}
	return contador
}

func main() {
	fmt.Println(procesarNumeros("Fizz", "Buzz"))
}

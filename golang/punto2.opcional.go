package main

import (
	"fmt"
	"math"
)

func esPrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func fibonacci(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

func main() {
	// Números entre 10 y 55, pares, excluyendo 16 y múltiplos de 3
	for num := 10; num <= 55; num++ {
		if num%2 == 0 && num != 16 && num%3 != 0 {
			fmt.Println(num)
		}
	}

	// Números primos del 1 al 1000
	primos := []int{}
	for num := 1; num <= 1000; num++ {
		if esPrimo(num) {
			primos = append(primos, num)
		}
	}
	fmt.Println(primos)

	// Serie de Fibonacci
	fibonacci(15)
}

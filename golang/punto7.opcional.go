// Función recursiva para calcular el factorial de un número
package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci8(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci8(n-1) + fibonacci8(n-2)
}

func main() {
	fmt.Println(factorial(5)) // Ejemplo: 5! = 120

	fmt.Println(fibonacci8(10)) // Ejemplo: Fib(10) = 55
}

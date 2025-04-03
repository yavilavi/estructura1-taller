package main

import (
	"fmt"
)

// Intercambio por valor
func intercambioPorValor(a, b int) (int, int) {
	temp := a
	a = b
	b = temp
	return a, b
}

// Intercambio por referencia
func intercambioPorReferencia(arr *[]int) {
	temp := (*arr)[0]
	(*arr)[0] = (*arr)[1]
	(*arr)[1] = temp
}

func main() {
	val1, val2 := 5, 10
	nuevosVal1, nuevosVal2 := intercambioPorValor(val1, val2)
	fmt.Println("Originales:", val1, val2)         // 5, 10
	fmt.Println("Nuevos:", nuevosVal1, nuevosVal2) // 10, 5

	refArr := []int{15, 20}
	intercambioPorReferencia(&refArr)
	fmt.Println("Referencia:", refArr[0], refArr[1]) // 20, 15
}

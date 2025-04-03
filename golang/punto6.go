package main

import (
	"fmt"
)

// Paso por valor
func modificarValor(num int) {
	num = 20
}

// Paso por referencia
func modificarReferencia(arr *[]int) {
	(*arr)[0] = 10
}

func main() {
	// Ejemplo por valor
	num1 := 10
	num2 := num1 // Se crea una copia
	num2 = 20
	fmt.Println("num1:", num1) // num1 sigue siendo 10
	fmt.Println("num2:", num2) // num2 ahora es 20

	str1 := "Hola"
	str2 := str1 // Se crea una copia
	str2 = "Mundo"
	fmt.Println("str1:", str1) // str1 sigue siendo "Hola"
	fmt.Println("str2:", str2) // str2 ahora es "Mundo"

	// Ejemplo por referencia
	arr1 := []int{1, 2, 3}
	arr2 := arr1 // Ambas variables apuntan al mismo slice
	arr2[0] = 10
	fmt.Println("arr1:", arr1) // arr1 también se modifica: [10, 2, 3]
	fmt.Println("arr2:", arr2) // arr2: [10, 2, 3]

	obj1 := map[string]int{"prop": 1}
	obj2 := obj1 // Ambas variables apuntan al mismo mapa
	obj2["prop"] = 2
	fmt.Println("obj1:", obj1) // obj1 también se modifica: {"prop": 2}
	fmt.Println("obj2:", obj2) // obj2: {"prop": 2}

	// Funciones por valor y referencia
	miNumero := 10
	modificarValor(miNumero)
	fmt.Println("miNumero:", miNumero) // Sigue siendo 10

	miArreglo := []int{1, 2, 3}
	modificarReferencia(&miArreglo)
	fmt.Println("miArreglo:", miArreglo) // miArreglo ahora es [10, 2, 3]
}

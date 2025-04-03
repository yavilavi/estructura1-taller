package main

import (
	"fmt"
	"math"
)

// Función sin parámetros ni retorno
func saludar() {
	fmt.Println("¡Hola, mundo!")
}

// Función con un parámetro y sin retorno
func saludarUsuario(nombre string) {
	fmt.Printf("¡Hola, %s!\n", nombre)
}

// Función con varios parámetros y con retorno
func sumar(a int, b int) int {
	return a + b
}

// Función dentro de una función
func externa() {
	fmt.Println("Función externa ejecutada")

	interna := func() {
		fmt.Println("Función interna ejecutada")
	}

	interna()
}

// Variables globales y locales
var globalVar = "Soy una variable global"

func pruebaVariables() {
	localVar := "Soy una variable local"
	fmt.Println(globalVar) // Accede a la variable global
	fmt.Println(localVar)  // Accede a la variable local
}

func main() {
	// Uso de funciones nativas del lenguaje
	fmt.Println(math.Max(10, 20))
}

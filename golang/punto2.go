package main

import (
	"fmt"
	"math"
)

func main() {
	// Operadores Aritméticos
	suma := 5 + 3
	resta := 10 - 4
	multiplicacion := 6 * 2
	division := 8 / 2
	modulo := 10 % 3
	exponente := math.Pow(2, 3)
	divisionEntera := 10 / 3

	// Operadores de Comparación
	esIgual := 5 == 5
	esDiferente := 5 != 3
	esMayor := 6 > 2
	esMenor := 2 < 8

	// Operadores Lógicos
	esVerdadero := true && false
	esFalso := true || false
	negacion := !true

	// Operadores de Asignación
	x := 10
	x += 5
	x -= 2

	// Operadores de Bits
	bitAnd := 5 & 3
	bitOr := 5 | 3
	bitXor := 5 ^ 3
	bitNot := ^5
	bitDesplazarIzq := 5 << 1
	bitDesplazarDer := 5 >> 1

	fmt.Println(suma, resta, multiplicacion, division, modulo, exponente, divisionEntera)
	fmt.Println(esIgual, esDiferente, esMayor, esMenor)
	fmt.Println(esVerdadero, esFalso, negacion)
	fmt.Println(x)
	fmt.Println(bitAnd, bitOr, bitXor, bitNot, bitDesplazarIzq, bitDesplazarDer)

	// Condicionales
	numero := 10
	if numero > 5 {
		fmt.Println("El número es mayor que 5")
	} else if numero == 5 {
		fmt.Println("El número es 5")
	} else {
		fmt.Println("El número es menor que 5")
	}

	// Bucles
	for i := 0; i < 5; i++ {
		fmt.Println("Iteración", i)
	}

	contador := 0
	for contador < 5 {
		fmt.Println("Contador:", contador)
		contador++
	}

	// Manejo de Excepciones (Go no tiene try/catch, se maneja con defer/recover)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error detectado:", r)
		}
	}()

	div := 10 / 0 // Esto causará un panic
	fmt.Println(div)

}

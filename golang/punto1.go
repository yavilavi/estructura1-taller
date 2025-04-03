package main

import "fmt"

// Crea un comentario en el código y coloca la URL del sitio web oficial del lenguaje de programación que has seleccionado.

// La página oficial es https://go.dev/

// Representa las diferentes sintaxis que existen de crear comentarios en el lenguaje (en una línea, varias...).

// Esta es la sintaxis de comentario de una línea

/*
   Esta es la sintaxis de comentario
   multilínea en este comentario puedo
   escribir varias líneas en un solo
   bloque de comentario
*/

// Crea una variable (y una constante si el lenguaje lo soporta).

func variable() {
	// Esto es una variable en Go
	var variable1 string = "Un valor tipo cadena"

	// Esto es una constante en Go
	const GRAVEDAD float64 = 9.81

	fmt.Println(variable1, GRAVEDAD)
}

// Crea variables representando todos los tipos de datos primitivos del lenguaje (cadenas de texto, enteros, booleanos...).

func tiposDatos() {
	// Entero (int)
	var numeroEntero int = 42

	// Flotante (float64)
	var numeroDecimal float64 = 3.14

	// Booleano (bool)
	var esMayor bool = true
	var esMenor bool = false

	// Cadena de texto (string)
	var mensaje string = "Hola, Go"

	fmt.Println(numeroEntero, numeroDecimal, esMayor, esMenor, mensaje)
}

// Imprime por terminal el texto: "¡Hola, [y el nombre de tu lenguaje]!"

func holaMundo() {
	fmt.Println("¡Hola, Go!")
}

// ejecutar la función holaMundo
func main() {
	holaMundo()
}

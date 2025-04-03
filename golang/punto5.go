package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "Hello, Go!"

	// Acceso a caracteres específicos
	fmt.Println(string(str[0])) // "H"

	// Longitud
	fmt.Println(len(str))

	// Subcadenas
	fmt.Println(str[0:5]) // "Hello"

	// Concatenación
	fmt.Println(str + " is awesome!")

	// Repetición (no hay operador directo, se usa strings.Repeat)
	fmt.Println(strings.Repeat("Ha! ", 3))

	// Recorrido
	for _, char := range str {
		fmt.Println(string(char))
	}

	// Conversión a mayúsculas y minúsculas
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))

	// Reemplazo
	fmt.Println(strings.ReplaceAll(str, "Go", "JavaScript"))

	// División y unión
	words := strings.Split(str, " ")
	fmt.Println(words)
	fmt.Println(strings.Join(words, "-"))

	// Interpolación
	name := "John"
	fmt.Printf("Hello, %s!\n", name)

	// Verificaciones
	fmt.Println(strings.Contains(str, "Go"))     // true
	fmt.Println(strings.HasPrefix(str, "Hello")) // true
	fmt.Println(strings.HasSuffix(str, "!"))     // true
}

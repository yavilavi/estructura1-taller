// Función recursiva que imprime números del 100 al 0
package main

import "fmt"

func printNumbers(num int) {
	if num < 0 {
		return
	}
	fmt.Println(num)
	printNumbers(num - 1)
}

func main() {
	printNumbers(100)
}

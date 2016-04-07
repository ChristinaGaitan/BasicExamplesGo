//Variables
package main

import "fmt"

func main() {
	
	// Si no se utiliza una variable si marca error el compilador
	// igual que con los imports no utilizados

	// Declaracian de variable
	var a string = "empezar"
	fmt.Println(a)

	// Declaraion de multiples variables
	var b, c int = 4, 5
	fmt.Println(b, c)

	// Declaración de variables asignandole un tipo
	var d = true
	fmt.Println(d)

	var d2 = "hola"
	fmt.Println(d2)

	// Declaracion de variable sin asignacion
	var e int
	fmt.Println(e)

	// Declaracion de variable sin utilizar la palabra reservada var
	f:= "short"
	fmt.Println(f)

	g:= 5+3
	fmt.Println(g)
}
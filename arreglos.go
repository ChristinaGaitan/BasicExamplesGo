// Tipos compuestos: Arreglos

package main

import "fmt"

func main() {
	// Declaracion de array
	var nombres[5]string

	// Declaración de array pre-llenado
	amigos := [5]string{"Raquel","Isabel","Fernando","Enrique","José"}

	nombres = amigos

	fmt.Println(nombres)

	// for i, nombres := range nombres{
	//  	fmt.Println(nombres, &nombres[i])
	//  }
}
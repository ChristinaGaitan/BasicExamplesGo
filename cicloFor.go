//Ciclo for

package main

import "fmt"

func main(){

	// Ciclo for funcionando como while
	fmt.Println("Alumnos")
	alumnos := 1
	for alumnos <= 3 {
		fmt.Println(alumnos)
		alumnos = alumnos + 1
	}

	// Ciclo for "normal"
	fmt.Println("Calificaciones")
	for calificaciones := 7;
		calificaciones <= 9;
		calificaciones++ {
		fmt.Println(calificaciones)
	}
}
// Métodos

package main

import "fmt"

type jugador struct {
	nombre string
	goles int
	partidos int
}

func (j *jugador) average()float64{
	return float64(j.partidos) / float64(j.goles)
}

func main(){
	jugadores := []jugador {
		{"carlos", 20, 60},
		{"fernando", 17, 30},
		{"alonso", 34, 60},
	}

	fmt.Println("Pending: print the average", jugadores)
	// for _, juador := range jugadores {
		//Pending print the average
		// fmt.Printf("Promedio: %f", average(&jugador{}))
	// }
}
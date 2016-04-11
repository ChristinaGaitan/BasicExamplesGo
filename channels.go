// Channels (canales)

package main

import (
	"fmt"
	"sync"
)

const (
	goroutines = 20
	capacidad = 4
)

//WaitGroup se usa para esperar que el programa termine
var wg sync.WaitGroup

// Recursos es un buffer channel para manejar los strings
var recursos = make(chan string, capacidad)

func main(){
	//Agregar el número de goroutines al GaitGroup
	wg.Add(goroutines)

	// Lanzar las goroutines necesarias para manejar el trabajo
	// Asegurarnos de poner la llamada para saber que las goroutines terminaron
	for gr := 1; gr <= goroutines; gr++ {
		go func(gr int){
			worker(gr)
			wg.Done()
		}(gr)
	}

	// Añadir los strings
	for s := 'A'; s < 'A'+capacidad; s++{
		recursos <- string(s)
	}

	wg.Wait()
}

//Lanzamos worker como una goroutine que procesa el trabajo del buffer channel
func worker(worker int){
	// Recibir string del channel
	valor :=  <- recursos

	//Imprimir el valor
	fmt.Printf("Worker: %d : %s\n", worker, valor)

	//Poner el string de regreso
	recursos <- valor
}
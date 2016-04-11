// Goroutines (threads)

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	// Alojar/reservar (allocate) un procesador lógico para que lo use el scheduler
	runtime.GOMAXPROCS(1)
}

func main(){
	// Declarar wait group para iniciar el conteo en las 2 goroutines
	var wg sync.WaitGroup
	wg.Add(2)
	
	fmt.Println("Iniciar Goroutines")

	// Declarar función anónima y crear una goroutine
	go func() {
		//Cuenta regresiva del 100 al 0
		for count := 100; count>=0; count-- {
			fmt.Printf("[A:%d]\n", count)
		}

		// Avisarle al main que ya se terminó
		wg.Done()
	}()


	// Declarar otra función anónima y crear otra goroutine
	go func() {
		//Cuenta regresiva del 100 al 0
		for count := 0; count<=100; count++ {
			fmt.Printf("[B:%d]\n", count)
		}
		
		// Avisarle al main que ya se terminó
		wg.Done()
	}()

	// Esperar a que terminen las goroutines
	fmt.Println("Espoerando a que terminen las goroutines")
	wg.Wait()

	fmt.Println("\n Finalizar el programa")
}
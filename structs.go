// Structs

package main

import "fmt"

type usuario struct{
	nombre string
	direccion string
	edad int
}

func main(){
	christina := usuario{
		nombre: "Christina",
		direccion: "Colima",
		edad: 30,
	}

	fmt.Println("Nombre: ", christina.nombre)
	fmt.Println("Dirección: ", christina.direccion)
	fmt.Println("Edad: ", christina.edad)

	// Struct anónimo
	nicole := struct{
		nombre string
		direccion string
		edad int
	} {
		nombre: "Nicole",
		direccion: "Tecomán",
		edad: 20,
	}

	fmt.Println("Nombre: ", nicole.nombre)
	fmt.Println("Dirección: ", nicole.direccion)
	fmt.Println("Edad: ", nicole.edad)
}
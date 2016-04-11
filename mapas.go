// Tipos compuestos: Mapas,
// equivalentes a hash tables o diccionarios

package main

import "fmt"

func main() {
	// Declaración de un mapa
	// La llave es de tipo string, los valores son de tipo int
	departamentos := make(map[string]int)

	// Asignación de datos
	departamentos["Devs"] = 25
	departamentos["Marketing"] = 50
	departamentos["Ejecutivos"] = 4
	departamentos["Ventas"] = 60
	departamentos["Mantenimiento"] = 8

	// Iterar los valores
	for key, value := range departamentos {
		fmt.Printf("Dept: %s Personas: %d\n", key, value)
	}
}
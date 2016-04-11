// Interfaces

package main

import "fmt"

type speaker interface{
	speak()
}

type ingles struct{}
type chino struct{}

// Implementación de la interface speaker
func (ingles) speak() {
	fmt.Println("Hello World")
}
func (chino) speak() {
	fmt.Println("你好")
}

func main(){
	var sp speaker

	// Asignarle un valor a la interface y llamaral método de la interface
	var i ingles
	sp = i
	sp.speak()

	var c chino
	sp = c
	sp.speak()

	// Crear nuevos valores y llamar la funcion
	decirHola(new(ingles))
	decirHola(&chino{})
}

// decirHola abstrae la función de hablar
func decirHola(sp speaker) {
	sp.speak()
}
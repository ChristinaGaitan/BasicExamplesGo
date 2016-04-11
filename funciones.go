// Funciones

package main

import "fmt"

// Struct usuario
type usuario struct{
	nombre string
	email string
}


// Crear un usuario (regresa apuntadores de valores de tipo usuario)
func nuevoUsuario() (*usuario, error) {
 return &usuario{"Christina", "christina@email.com"}, nil
}

func main() {
	usuario1, error1 := nuevoUsuario()
	if error1 != nil {
		fmt.Println(error1)
		return
	}
	fmt.Println(*usuario1)
	
	_, error2 := nuevoUsuario()
	if error2!= nil {
		fmt.Println(error2)
		return
	}
}
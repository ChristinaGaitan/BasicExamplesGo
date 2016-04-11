// Range

package main

import "fmt"

func main(){
	numeros := []int{2,4,6}
	suma := 0
	for _, numero := range numeros {
		suma += numero
	}

	fmt.Println("Suma: ", suma)

	for i, numero := range numeros{
		fmt.Println("Index: ", i, numero)
	}

	algo := map[string]string{"a":"auto", "b":"bebé"}
	for key, value := range algo {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}
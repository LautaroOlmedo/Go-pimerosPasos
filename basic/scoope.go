package basic

import "fmt"

// ALCANCE DE LA VARIABLE
var scoope0 int // ---> Scoope global
func Scope(){
	scoope1 := 1 // ---> Scoope dentro de la función 

	{
		scoope2 := new([]int) // ---> Scoope dentro de {}
		fmt.Printf("El tipo de %v es (%T) \n", scoope2, scoope2)
	}

	fmt.Printf("El tipo de %v es (%T) \n", scoope0, scoope0)
	fmt.Printf("El tipo de %v es (%T) \n", scoope1, scoope1)

}
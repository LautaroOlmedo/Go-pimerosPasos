package basic

import (
	"fmt"
	"strconv"
)

func Variables(){
	var( // FORMAS DE DECLARAR VARIABLES
		name string
		age uint8
		dni string
		pensioner bool
		pensionerStr string // Creamos 2 variables ya que no podemos ingresar por consola datos booleanos entonces lo ingresamos como string y dps lo parseamos
	)
	// var myStr string // FORMAS DE DECLARAR VARIABLES
	// strVar := "Lautaro" // FORMAS DE DECLARAR VARIABLES

	fmt.Println("Entry your name")
	fmt.Scan("%s", &name)

	fmt.Println("Entry your age")
	fmt.Scan("%d", &age)

	fmt.Println("Entry your DNI")
	fmt.Scan("%s", &dni)

	fmt.Println("Are you pensioner?")
	fmt.Scan("%s", &pensionerStr)

	pensioner, _ = strconv.ParseBool(pensionerStr) // De string a bool. _ hace referencia a que ignoremos el posible error

	fmt.Printf("Your name is %v (%T) \n", name, name)
	fmt.Printf("Your age is %v (%T) \n", age, age)
	fmt.Printf("Your dni is %v (%T) \n", dni, dni)
	fmt.Printf("Is pensioner: %v (%T) \n", pensioner, pensioner)
}
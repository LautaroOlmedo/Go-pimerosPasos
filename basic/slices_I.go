package basic

import "fmt"

func SlicesI() {
	var marcaAutos = make([]string, 0) // Valor inicial ---> 0
	 // marcaAutos := []string{} // Valor inicial ---> 0

	fmt.Println("Ingrese una marca de autos")

	for{

		var marcaCapturada string
		fmt.Scanln(&marcaCapturada)

		if(marcaCapturada == ""){
			break
		}else{
			marcaAutos = append(marcaAutos, marcaCapturada) // Agregar valores "append"
		}
	}
	fmt.Println("Las marcas de autos son: ", marcaAutos)
}
package basic

import "fmt"

func ArrayII() {
	var(
		myArr [4]string
	)
	//marcasAutos := [...]string{"Toyota", "Peugeot", "Ferraro"} // ---> Podemos ir agregando cuantos valores querramos dentro de {}

	myArr[0] = "Lautaro"
	myArr[1] = "Fran"
	myArr[2] = "Cami"
	myArr[3] = "Juli"
	fmt.Print(myArr)

	for k, v := range myArr { // ---> Gracias a "range" obtendremos en "k" el valor de la posición(0, 1, 2, 3) y en "v" el valor de cada posición
		fmt.Println(k, v)
	}


}

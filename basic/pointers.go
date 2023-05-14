package basic

import "fmt"

// PUNTEROS. UTILES CUANDO REALIZAMOS CALCULOS CON VARIABLES. PASAMOS SU REFERENCIA Y NO CREAMOS UNA COPIA. AUMENTA LA EFICIENCIA DE NUESTROS PROGRAMAS

func Pointers(){
	auto1 := "Pirus"
	auto2 := "Tesla"

	fmt.Println("auto 1: ", auto1, &auto1) // & HACE REFERENCIA A SU DIRECCIÓN EN MEMORIA
	fmt.Println("auto 2: ", auto2, &auto2)

	auto3 := auto1 // VALOR POR VALOR
	fmt.Println("auto 3: ", auto3)

	auto4 := &auto2 // VALOR POR REFERENCIA (Dirección)
	fmt.Println("auto 4: ", auto4, *auto4) // *auto4 es un puntero. Dame su VALOR
}
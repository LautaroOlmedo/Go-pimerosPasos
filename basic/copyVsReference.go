package basic

import "fmt"

// ENVIAR VALORES POR COPIA O POR PUNTEROS.

func CopyVsReference(){
	var stature float32
	fmt.Println("Entry your stature in mts")
	fmt.Scanf("%f", &stature)
	fmt.Println(&stature) // ---> Devolver dirección en memoria. Dirección1

	// stature = mtsInFeets(stature) // ---> Por copia
	mtsInFeets2(&stature) // Por puntero
	fmt.Println("Your stature in feets is", stature)
}

// func mtsInFeets(stature float32) float32{ // ---> Por copia
// 	fmt.Println("stature, COPY", &stature) // Dirección2. Será disinta de Dirección1 ya que copia el valor y crea un nuevo espacio
// 	stature = stature * 3.28
// 	return stature
// }

func mtsInFeets2(stature *float32){ // ---> Por puntero
	fmt.Println("stature, COPY", stature) // Dirección2. No lleva & ya que es un puntero. Dirección2. Será igual a Dirección1 ya que recibe el mismo espacio en memoria gracias al puntero. No tiene return porque es un puntero. Modifica ese espacio en memoria
	*stature = *stature * 3.28
}
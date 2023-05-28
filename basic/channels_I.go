package basic

import (
	"fmt"
	"time"
)

// PERMITEN A LOS DATOS MOVERSE DENTRO YN FUERA DE LAS RUTINAS, FACILTAN LA COMUNICACIÓN ENTRE ELLAS
// "no comunicarse compartiendo la memoria, compartir la memoria al comunicarse". NO PODEMOS TENER MAS DE UNA RESOLUCIÓN POR LECTURA

func lecture() string{
	time.Sleep(time.Second * 2)
	return "this is te content of the file"

}

func Channels(){
	myChannel := make(chan string) // ---> Creación de un canal

	go func(){ // ---> Runtina
		myChannel <- lecture() // ---> El bloqueo se resuelve acá
	}()
	fmt.Println(<- myChannel) // ---> SE GENERA UN BLOQUEO ACÁ (lectura). y espera a que chanel tenga un valor
}
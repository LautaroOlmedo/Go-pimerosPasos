package basic

import (
	"fmt"
	"time"
)

// CONSISTE EN INVOCAR UNA FUNCIÓN QUE REQUIERE X CANTIDAD DE TIEMPO PARA QUE LA FUNCIÓN SE COMPLETE. LAS FUNCIONES QUE REQUIEREN TIEMPO
// SUELEN LLAMAR A SERVICIOS DE UN 3RO (API), LECTURA DE ARCHIVOS GRANDE, CONECTAN CON LA DDBB, ENVIAN EMAILS, ETC...

func FunctionsAndBlocking(){
	downloadFiles()
	fmt.Println("FINISH")

}

func downloadFiles(){
	// for i := 0; i < 1e6; i++ {
	// 	fmt.Printf("\r Downloading %v MB of 1e6 MB", i)
	// }
	time.Sleep(time.Second * 5) // ---> Al ejecutar time.Sleep se detiene la ejecución durante el perdiodo de tiempo asignado. Genera bloaqueo a través de la función
	fmt.Println("Download")
}
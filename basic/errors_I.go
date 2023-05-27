package basic

// GO UTILIZA UN CONCEPTO DISTINTO EN CUANTO AL MANEJO DE ERRORES. NO UTILIZA TRY CATCH. LOS ERRORES SON UN "TYPE" Y PUEDEN SER ENVIADOS A TRAVÉS DE FUNCIONES
// Y MÉTODOS. GO RETORNA UN ERROR COMO ÚLTIMO PARÁMETRO POSIBLE

// LOS ERRORES DEBEN ESPECIFICAR EL PROBLEMA, OFRECER UNA SOLUCIÓN, SER RESPETUOSO CON EL USUARIO

import (
	"fmt"
	"io/ioutil"
)

func ErrorsI(){
	//text, err := ioutil.ReadFile("archivo.txt") // ---> err. Ultimo parámetro
	text, err := readFiles("archivo.txt")
	if(err != nil){
		fmt.Println(err)
	}else{
		fmt.Println(string(text))
	}
}

func readFiles(file string) (textFile []byte, err error){
	return ioutil.ReadFile(file)
}
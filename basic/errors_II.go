package basic

import (
	"errors"
	"fmt"
	"io/ioutil"
)

// CREAR ERRORES PERSONALIZADOS. ERROR TYPE

func readFilesII(file string) (textFile []byte, err error){ // ---> Se llama así porque re creó previamente en errors_I una función con ese nombre
	textFile, err = ioutil.ReadFile(file) // ---> Cuando utilizamos reasignación de variables pasadas por marámetros no utilizamos ":=". Utilizamos "="
	if err != nil {
		err = errors.New(("No se puede leer el archivo" + file)) // ---> errors.New nos permite sobrescribir nuestro error personalizado
	}
	return
}

func ErrorsII(){
	text, err := readFilesII("archivo.txt")

	if(err != nil){
		fmt.Println(err)
	}else{
		fmt.Println(text)
	}
}
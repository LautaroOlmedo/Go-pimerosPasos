package basic

import (
	"fmt"
	"reflect" // Con reflect podes acceder al tipo de dato
)


func FirstArray(){

	var alumnos[3]string

  for i := 0; i < 3; i++ {
	fmt.Println("Ingrese el nombre para el alumno", i)
	fmt.Scanf("%s", &alumnos[i])
  }

  for t := 0; t < 3; t++ {
	fmt.Println("El nombre del alumno", t, "es", alumnos[t])
	fmt.Println("El tipo de dato es", reflect.TypeOf(alumnos[t]))
	
  }

}
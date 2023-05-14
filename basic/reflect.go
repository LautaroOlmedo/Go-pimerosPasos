package basic

import (
	"fmt"
	"reflect" // REFLECT VERIFICA TIPOS DE DATOS
)

func Reflect(){
	var student string = "Lautaro"
	var age int = 23
	var weight float32 = 54.23

	fmt.Println("The variable student its a ", reflect.TypeOf(student), " type, and the variable's value is", student)
	fmt.Println("The variable age its a ", reflect.TypeOf(age), " type, and the variable's value is", age)
	fmt.Println("The variable studnets its a ", reflect.TypeOf(weight), " type, and the variable's value is", weight)

	fmt.Printf("The variable studnets its a %T type, and the variable's value is %v \n", student, student)
	fmt.Printf("The variable studnets its a %T type, and the variable's value is %v \n", age, age)
	fmt.Printf("The variable studnets its a %T type, and the variable's value is %v \n", weight, weight)

}
package basic

import "fmt"

// CON EL PAQUETE FMT PODEMOS CREAR ERRORES Y UTILIZARLOS CON COMODINES

func ErrorsIII(){

	age := 18
    
	// nerr = errors.New(("No se puede leer el archivo" + file)) // ---> Sin "fmt". Necesita librería "errors"
	// err := fmt.Errorf("Age is less than %v", age) // ---> Con "fmt". No necesita librería "errors"

	if err := valdiateAge(age) ; err != nil{
		fmt.Println(err)
	}
}

func valdiateAge(age int) error{
	return fmt.Errorf("Age is less than %v", age)
}
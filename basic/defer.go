package basic

import "fmt"

// APLAZAMIENTO DE FUNCIONES. Se utiliza generalmente para tareas de "limpieza" como cerrar conexiones con BBDD, borrar archivos temporales, etc...
// DEFER TIENE EL COMPORTAMIENTO LAST IN FIRST OUT

func Defer(){
	connect()

	// deleteCache:= func ()  {
	// 	fmt.Println("removing cache from the system")
	// }
	defer func ()  {
		fmt.Println("removing cache from the system")
	}()
	// defer deleteCache()
	defer deleteLogs()
	defer disconect() // ---> Ejecuta la función al final de la ejecución de todo lo que resta del programa, independientemente de lo que suceda
	read()
	write()
	
}

func connect(){
	fmt.Println("Connected to BBDD")
}

func disconect(){
	fmt.Println("Disconnected from BBDD")
}

func read(){
	fmt.Println("Reading data from BBDD")
}

func write(){
	fmt.Println("Writing on BBDD")
}

func deleteLogs(){
	fmt.Println("removing records from the system")
}
package basic

import (
	"fmt"
	"time"
)

func StringsII(){
	// \n, \t, \r, \

	// \t ---> Se utiliza para crear tablas
	// \r ---> Retorna la impresión al inicio de la linea. Devuelve el cursor al principio

	// ---------- ---------- ---------- ---------- ----------

	// fmt.Println("Student\t\t\tCalification")
	// fmt.Print("Lautaro\t\t\t95\n")
	// fmt.Print("Julian\t\t\t100\n")

	// ---------- ---------- ---------- ---------- ----------

	// fmt.Printf("Hola Mundo")
	// fmt.Printf("\r")
	// fmt.Printf("Adios Mundo")
	// fmt.Print("\n")
	// fmt.Print(" \"n\" ")
	// %% //

	// ---------- ---------- ---------- ---------- ----------
	var megas int
	var file string

	fmt.Printf("Entry the file's name")
	fmt.Scanln(&file)

	fmt.Printf("How many megas use the file?")
	fmt.Scanln(&megas)

	downloandTorrent(file, megas)
}

func downloandTorrent(file string, megas int){
	
	for i:= 1; i <= megas; i++{
		rate := float64(i) / float64(megas) * 100 
		fmt.Printf("Downloading %v, %v to %v (%.0f%%)\r", file, i, megas, rate)
		time.Sleep(time.Microsecond * 10)

	}
	fmt.Printf("\n")
	fmt.Println("File downloaded successfully")

}
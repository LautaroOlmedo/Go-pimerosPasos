package basic

// RUTINAS: FORMA DE TRABAJAR CON PROGRAMACIÓN CONCURRENTE
import (
	"fmt"
	"time"
)

func viceHours(){
	time.Sleep(time.Second * 5)
	fmt.Println("Lautaro plays 0 hours per day")
	fmt.Println("Francisco plays 10 hours per day. He's a rat")
}

func Routines(){
	go viceHours() // ---> Pone a la función en concurrencia
	time.Sleep(time.Second * 6)
	fmt.Println("END")
}
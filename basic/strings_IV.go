package basic

import "fmt"

// CONVERSIONES Y TIPOS PERSONALIZADOS

type Num int

func (n Num) toString() string {
	return fmt.Sprint(n)
}

func (n Num) elevate() Num {
	return n * n
}

func (n Num) haft() float64 {
	return float64(n) / 2
}

func StringsIV(){
	var students Num = 10

	fmt.Println(students.toString())
	fmt.Println(students.elevate())
	fmt.Println(students.haft())

}
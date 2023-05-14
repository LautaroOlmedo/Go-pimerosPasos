package basic

import "fmt"

// FUNCIONES ANÓNIMAS/LAMBDAS

func FunctionsV(){
	hello := func ()  {
		fmt.Println("Hello World!!")
	}
	hello()
}

func FunctionsV2(){
	name := "Lautaro"
	hello := func ()  {
		fmt.Println("Hello World!! My name is", name)
	}
	hello()
}

func FunctionsV3(radius float64) (area func() float64, perimeter func() float64){ // ---> area, perimeter := FunctionsV3(7.09). Para invocar area(), perimeter()
	const Pi = 3.1416
	area = func() float64 {
		return Pi * (radius * radius)
	}
	perimeter = func() float64 {
		return 2 * Pi * radius
	}
	return
}
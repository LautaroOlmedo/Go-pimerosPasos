package basic

import "fmt"

type Caracteristicas struct{
	doors uint16
	wheels uint16
}

type Car struct{
	Model string
	Brand string
	Description Caracteristicas
}



func StructIV(){

	p208 := Car{}
	p208.Model = "Peugeot 208"
	p208.Brand = "Peugeot"
	p208.Description = Caracteristicas{}
	p208.Description.doors = 4
	p208.Description.wheels = 4


	
	p208.printCarInfo()


}

func (des Caracteristicas) printCarDescription(){
	fmt.Printf("Puertas: %+v . Ruedas: %+v \n", des.doors, des.wheels)
}

func (car Car) printCarInfo(){
	
	fmt.Printf("Car: %+v \n", car.Model)
	fmt.Printf("Brand: %+v \n", car.Brand)
	car.Description.printCarDescription()
}
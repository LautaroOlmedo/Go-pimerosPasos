package basic

import "fmt"

type Dog struct{
	Raze string
}

func StructIII(){
	dogo := Dog{
		Raze: "Dogo",
	}

	sanBernardo := new(Dog) // ---> Puntero
	caniche := new(Dog)

	sanBernardo.Raze = "San Bernardo"
	caniche.Raze = "Caniche"

	// fmt.Printf("%+v %p \n", dogo, &dogo)
	// fmt.Printf("%+v %p \n", sanBernardo, sanBernardo)
	// fmt.Printf("%+v %p \n", caniche, caniche)
	PrintDogInfo(&dogo)
	PrintDogInfo(sanBernardo)
	PrintDogInfo(caniche)

	

	

}
func PrintDogInfo(d *Dog){
	fmt.Printf("%+v %p \n",d, d)
}

package basic

// FUNCIONES QUE RETORNAN ERRORES

import "fmt"

func sellBeer(age int, quantity int) (int, error){
	if(age < 18){
		return 0, fmt.Errorf("You cannot drink if you are under 18 years old. You have %v", age)
	}
	return quantity, nil
}

func ErrorsIV(){

	var age, quantity int

	fmt.Println("Enter your age")
	fmt.Scanln(&age)

	fmt.Println("How many beers do u want?")
	fmt.Scanln(&quantity)

	if _, err := sellBeer(age, quantity) ; err != nil{
		fmt.Println(err)
	}

	fmt.Printf("you bought %v beers", quantity)
}

// FUNCIONES QUE RETORNAN ERRORES
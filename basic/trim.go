package basic

import (
	"fmt"
	"strconv"
	"strings"
)

// ELIMINAR ESPACIOS EN BLANCO

type User struct{
	Name string
	Age uint8
	DNI string
	Nacionality string
}



func AddUser(user *User) {
	fmt.Printf("Enter the user name (%v)\n", user.Name)
	fmt.Printf("Enter the user age (%v)\n", user.Age)
	fmt.Printf("Enter the user DNI (%v)\n", user.DNI)
	fmt.Printf("Enter the user nacionality (%v)\n", user.Nacionality)

	
}

func RegisterUser() {
	data := map[string]string{
		"name": "  Jose",
		"age": "23  ",
		"DNI": "42509723",
		"Nacionality": "  Argentine  ",
	}
	
	name := strings.TrimLeft(data["name"], " ") // ---> Elimina los espacios del lado izquierdo
	age, err := strconv.Atoi(strings.TrimRight(data["age"], " ")) // ---> Convierte de string a INT y elimina los espacios del lado derecho
	if (err != nil){
		fmt.Println("Wrong age")
		return 
	}
	DNI := data["DNI"]
	nacionality := strings.Trim(data["Nacionality"], " ") // ---> Elimina los espacios del lado izquierdo y derecho
	
	newUser := new(User)
	newUser.Name = name
	newUser.Age = uint8(age)
	newUser.DNI = DNI
	newUser.Nacionality = nacionality
	AddUser(newUser)
}



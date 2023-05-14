package main

import (
	"firstSteps/basic"
	"fmt"
	//"fmt"
)

func main(){
	employed1 := basic.CreateEmployed("Lautaro", 23, 3500.50, "Software Developer", "argentine")
	fmt.Println(employed1.GetSalary())

	fmt.Printf("%+v \n", employed1)
}
package basic

import (
	"fmt"
	"strconv" // CONVERSIÓN DE DATOS
)

func StrConv(){
	var isAdultString string = "true"
	isAdultBool, _  := strconv.ParseBool(isAdultString) // De string a bool
	toString := strconv.FormatBool(false) // de Bool a string
	fmt.Printf("is adult %v and isAdultBool type is %T \n", isAdultBool, isAdultBool)
	fmt.Printf("toString is a %v  variable. toString type is %T \n", toString, toString)
}
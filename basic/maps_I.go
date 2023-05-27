package basic

import "fmt"

func MapsI(){
	var daysOfWeek = make(map[string]uint8) // ---> Los maps no están ordenados

	daysOfWeek["Mondai"] = 1
	daysOfWeek["Tuesday"] = 2
	daysOfWeek["Wednesday"] = 3
	daysOfWeek["Thursday"] = 4
	daysOfWeek["Friday"] = 5

	for k, v := range daysOfWeek {
		fmt.Printf("%v is the %v of the week \n", k, v )
	}

	fmt.Println(" ----- ")

	countries := map[string]string{ // ---> Otra forma de crear maps
		"Mexico": "CDMX",
		"Argentina": "Buenos Aires",
		"Italia": "Roma",
	}

	for country, capital := range countries {
		fmt.Printf("The capital of %v is %v \n", country, capital )
	}
}
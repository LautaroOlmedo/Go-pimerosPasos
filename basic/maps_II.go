package basic

import "fmt"

func MapsII(){
	dogs := map[string]int{
		"salchicha": 1,
		"dogo": 2,
		"caniche": 3,
		"golden": 4,
		
	}

	delete(dogs, "salchicha") // ---> Eliminar elementos de un map
	fmt.Print(dogs)
}
package basic

import "fmt"

func SLicesII(){
	// Eliminar elementos de un Slice. NO podemos eliminar directamente, en Go REASIGNAMOS
	razasPerro := []string{"Labrador", "Golden", "Pug", "Caniche", "Beagle"}
	razasPerro = append([]string{}, razasPerro[0], razasPerro[1], razasPerro[2], razasPerro[4]) // ---> REASIGNACIÓN de forma manual. Se creó un nuevo Slice sin "caniche"
	// sliceIzq := razasPerro[0:2] // ---> Creación a partir de otro Slice. Se creó un nuevo Slice con los valores "Labrador" y "Golden" ya que la ultima pos. no es inclusiva
	// sliceDer := rarazasPerro[2:] || := rarazasPerro[:2] // ---> De esta forma la ultima o primera posición si estaría incluida
	fmt.Println(razasPerro)
}
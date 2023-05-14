package basic

// FUNCIONES CON MULTIPLES VALORES DE RETORNO

func FunctionsII(radius float64) (area float64, perimeter float64){
	const Pi = 3.14
	area =  Pi * radius * radius
	perimeter = 2 * Pi * radius
	return 

}

// area, permimetro := FunctionsII(5.10) ---> Así sería el llamado. Ya que retorna 2 valores creamos dos variables para almacenar los resultados
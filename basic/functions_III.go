package basic

// FUNCIONES VARIÁDICAS. FUNCIONES QUE RECIBEN N NUMEROS DE PARAMETROS.

func FunctionsIII(nums... int) (sum int){
	for _, val := range nums{
		sum = sum + val
	}
	return
}

func FunctionsIII2(initialValue int, nums... int) (sum int){ // PODEMOS AGREGAR UN VALOR INICIAL 
	sum = initialValue
	for _, val := range nums{
		sum = sum + val
	}
	return
}
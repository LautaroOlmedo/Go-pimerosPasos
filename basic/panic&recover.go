package basic

import "fmt"

// PANIC CORTA LA EJECUCIÓN DEL PROGRAMA E INICIA UN PROCESO DE PÁNICO. NO ES RECOMENDADO UTILIZAR SU USO


func blockCreditCard(){
	if r := recover() ; r != nil{ // ---> Llama a recover SI HAY un bloqueo. Si recover devuelve un valor diferente de nil se puede recuperar el programa. Paso 2
		fmt.Println("Blocked credit card")
		fmt.Println("Get in touch with the customer")
	}
}
func PanicAndRecover(){

	defer blockCreditCard() // ---> Ejecuta al finalizar. Paso 3
	var moneyInBanc float64 = 1000
	var retiro float64

	for {
		
		fmt.Println("How many dollars do you want to withdraw?")
		fmt.Scanln(&retiro)

		if(moneyInBanc < retiro){
			panic("ABORT") // ---> Bloquea la ejecución. Paso 1
		}
		moneyInBanc -= retiro
		fmt.Printf("Your stract was aprobed. You withdrew $%v. Your available money is $%v", retiro, moneyInBanc)
	}
}
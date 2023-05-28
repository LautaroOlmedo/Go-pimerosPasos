package basic

import "fmt"

// ITERAR CANALES

func ChannelsII(){
	news := make(chan string, 5) // ---> chanel, type que contendrá, longitud

	news <- "Argentina"
	news <- "Francia"
	news <- "Marruecos"
	news <- "Croacia"
	//news <- "Chile NO"

	close(news) // ---> Al momento de poner close y cerrar el canal ya no podremos escribir dentro del canal
	
	i := 0
	for v := range news {
		fmt.Println(i, v)
		i++
	}
}
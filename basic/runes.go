package basic

import "fmt"

func Runes(){
	// string | []unit8 / []byte | ASCII
	// []int32 / []rune | utf8

	saludo := []rune("😏 🥰 🔥")

	for i := 0; i < len(saludo); i++{
		v := saludo[i]
		fmt.Printf("%v %T %c %x %U %b \n", v, v, v, v, v, v)

	}
}
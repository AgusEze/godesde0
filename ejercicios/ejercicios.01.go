package ejercicios

import (
	"strconv"
)

func ConvertiraNumero(texto string) (int, string) {
	//var numero int
	var palabra string
	numero, _ := strconv.Atoi(texto)
	//fmt.Println(numero)
	if numero > 100 {
		palabra = "Es mayor a 100"
	} else {
		palabra = "Es menor a 100"
	}

	//fmt.Println(palabra)
	return numero, palabra
}

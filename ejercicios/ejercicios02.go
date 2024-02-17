package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var multiplicar int
var err error

func IngreseNumero() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Por favor ingrese un numero entre el 1 y el 10: ")

		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())

			if err != nil || numero <= 0 || numero >= 11 {
				fmt.Println("El dato ingresado es erroneo, por favor vuelva a intentarlo")
			} else {
				break
			}

		}
	}
	if numero >= 1 && numero <= 10 {
		for i := 1; i < 11; i++ {
			multiplicar = numero * i
			fmt.Println("Tabla de multiplicar:", numero, "X", i, "=", multiplicar)
		}
	}

	//El dato que ingreso el usuario se guarda en numero

}

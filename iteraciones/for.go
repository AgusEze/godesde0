package iteraciones

import "fmt"

/*func Iterar() {				//Forma "larga"
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}*/

/*func Iterar() { //Forma "correcta"
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}*/

func Iterar() {
	for i := 10; i > 1; i-- {
		if i == 6 {
			break //BREAK Directamente termina con el programa en el mismo renglon
			//continue// solo se salta la linea actual
		}

		fmt.Println(i)
	}
}

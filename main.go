package main

import (
	"fmt"

	"github.com/AgusEze/godesde0/variables"
)

func main() {
	estado, texto := variables.ConvertiraTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

}

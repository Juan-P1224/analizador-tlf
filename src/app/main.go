package main

import (
	"fmt"

	"analizador/src/modelo"
)

func main() {

	an := modelo.NewAnalizadorLexico("892892asas")

	// Imprimir mensaje de depuración
	fmt.Println("Antes de Analizar")

	an.Analizar()

	// Imprimir mensaje de depuración
	fmt.Println("Después de Analizar")

	// Imprimir la lista de tokens
	for _, token := range an.GetListaTokens() {
		fmt.Println(token)
	}

}

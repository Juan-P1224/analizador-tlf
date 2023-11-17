package main

import (
	"fmt"

	"analizador/src/modelo"
)

// Mapa para mapear valores de Categoría a sus nombres
var categoriasNombres = map[modelo.Categoria]string{
	modelo.NO_RECONOCIDO:       "NO_RECONOCIDO",
	modelo.ENTERO:              "ENTERO",
	modelo.DECIMAL:             "DECIMAL",
	modelo.IDENTIFICADOR:       "IDENTIFICADOR",
	modelo.PALABRA_RESERVADA:   "PALABRA_RESERVADA",
	modelo.CADENA_CARACTERES:   "CADENA_CARACTERES",
	modelo.COMENTARIO_LINEA:    "COMENTARIO_LINEA",
	modelo.COMENTARIO_BLOQUE:   "COMENTARIO_BLOQUE",
	modelo.OPERADOR_ARITMETICO: "OPERADOR_ARITMETICO",
	modelo.OPERADOR_RELACIONAL: "OPERADOR_RELACIONAL",
	modelo.OPERADOR_LOGICO:     "OPERADOR_LOGICO",
	modelo.OPERADOR_INCREMENTO: "OPERADOR_INCREMENTO",
}

func main() {
	// Imprimir mensaje de depuración
	fmt.Println("Antes de Analizar")

	codigoFuente := `
	892892
	asas
	// Esto es un comentario
	float variable = 3.14
	"cadena de caracteres"
	== != <= >= < >
	`
	an := modelo.NewAnalizadorLexico(codigoFuente)
	an.Analizar()

	// Imprime la lista de tokens
	for _, token := range an.ListaTokens {
		// Obtener el nombre de la categoría desde el mapa
		nombreCategoria := categoriasNombres[token.Categoria]
		fmt.Printf("%s -> %s\n", token.Palabra, nombreCategoria)
	}
	// Imprimir mensaje de depuración
	fmt.Println("Después de Analizar")
}

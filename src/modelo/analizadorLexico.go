package modelo

import (
	"fmt"
	"unicode"
)

type AnalizadorLexico struct {
	CodigoFuente string
	ListaTokens  []Token
}

func NewAnalizadorLexico(codigoFuente string) *AnalizadorLexico {
	return &AnalizadorLexico{
		CodigoFuente: codigoFuente,
		ListaTokens:  make([]Token, 0),
	}
}

// Analizar extrae los tokens de un código fuente dado
func (al *AnalizadorLexico) Analizar() {
	// El primer token se extrae a partir de la posición cero
	//i := 0

	// TODO: Ciclo para extraer todos los tokens y los vaya guardando en la lista de tokens
}

// ExtraerSgteToken extrae el próximo token del código fuente a partir de la posición dada
func (al *AnalizadorLexico) ExtraerSgteToken(indice int) *Token {
	var token *Token

	// Intenta extraer un entero
	token = al.extraerEntero(indice)
	if token != nil {
		return token
	}

	// TODO: Llamar acá todos los métodos de extraer, extraerDecimal, extraerIdentificador, etc.

	// Extrae un token no reconocido
	token = al.extraerNoReconocido(indice)

	return token
}

// ExtraerEntero intenta extraer un entero del código fuente a partir de la posición dada
func (al *AnalizadorLexico) extraerEntero(indice int) *Token {
	if unicode.IsDigit(rune(al.CodigoFuente[indice])) {
		posicion := indice

		for indice < len(al.CodigoFuente) && unicode.IsDigit(rune(al.CodigoFuente[indice])) {
			indice++
		}

		return &Token{
			Palabra:         al.CodigoFuente[posicion:indice],
			Categoria:       ENTERO,
			IndiceSiguiente: indice,
		}
	}

	return nil
}

// ExtraerNoReconocido extrae un símbolo no reconocido del código fuente a partir de la posición dada
func (al *AnalizadorLexico) extraerNoReconocido(indice int) *Token {
	lexema := al.CodigoFuente[indice : indice+1]
	return &Token{
		Palabra:         lexema,
		Categoria:       NO_RECONOCIDO,
		IndiceSiguiente: indice + 1,
	}
}

// GetListaTokens obtiene la lista de todos los tokens reconocidos por el analizador léxico
func (al *AnalizadorLexico) GetListaTokens() []Token {
	return al.ListaTokens
}

func main() {
	// Ejemplo de uso del AnalizadorLexico
	codigoFuente := "123 + 456"
	analizador := NewAnalizadorLexico(codigoFuente)
	analizador.Analizar()

	// Imprimir la lista de tokens
	for _, token := range analizador.GetListaTokens() {
		fmt.Printf("Lexema: %s, Categoría: %d, Posición: %d\n", token.Palabra, token.Categoria, token.IndiceSiguiente)
	}
}

package modelo

import (
	"regexp"
	"strings"
	"unicode"
)

// AnalizadorLexico representa la clase principal del analizador léxico del compilador.
type AnalizadorLexico struct {
	CodigoFuente string
	ListaTokens  []*Token
}

// NewAnalizadorLexico crea una nueva instancia de AnalizadorLexico
func NewAnalizadorLexico(codigoFuente string) *AnalizadorLexico {
	return &AnalizadorLexico{
		CodigoFuente: codigoFuente,
		ListaTokens:  make([]*Token, 0),
	}
}

// Expresiones Regulares
var (
	regexIdentificador      = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
	regexNumNatural         = regexp.MustCompile(`^[0-9]+$`)
	regexNumReal            = regexp.MustCompile(`^[0-9]+(\.[0-9]+)?$`)
	regexPalabraReservada   = regexp.MustCompile(`^(if|else|while|int|float|boolean|return)$`)
	regexOperadores         = regexp.MustCompile(`^(\+|\-|\*|\/|==|!=|<=|>=|<|>|&&|\|\|)$`)
	regexCadena             = regexp.MustCompile(`^\"[^\"]*\"$`)
	regexComentario         = regexp.MustCompile(`^(\/\/[^\n]*|\/\*[\s\S]*?\*\/)$`)
	regexOperadorAsignacion = regexp.MustCompile(`^=$`)
	regexParentesisApertura = regexp.MustCompile(`^\($`)
	regexParentesisCierre   = regexp.MustCompile(`^\)$`)
	regexLlaveApertura      = regexp.MustCompile(`^{`)
	regexLlaveCierre        = regexp.MustCompile(`^}`)
	regexTerminal           = regexp.MustCompile(`^;$`)
	regexSeparador          = regexp.MustCompile(`^,$`)
	regexHexadecimal        = regexp.MustCompile(`^0x[0-9A-Fa-f]+$`)
)

// Función para extraer un token basado en la expresión regular proporcionada
func (al *AnalizadorLexico) extraerTokenConExpresionRegular(expReg *regexp.Regexp, indice int) *Token {
	subcadena := al.CodigoFuente[indice:]
	match := expReg.FindStringIndex(subcadena)

	if match != nil && match[0] == 0 {
		lexema := subcadena[:match[1]]
		return &Token{
			Palabra:         lexema,
			Categoria:       determinarCategoriaPorExpresionRegular(expReg),
			IndiceSiguiente: indice + len(lexema),
		}
	}

	return nil
}

// Función para determinar la categoría de un token basado en la expresión regular
func determinarCategoriaPorExpresionRegular(expReg *regexp.Regexp) Categoria {
	switch expReg {
	case regexIdentificador:
		return IDENTIFICADOR
	case regexNumNatural, regexNumReal:
		return ENTERO // O DECIMAL, depende de la lógica de tu lenguaje
	case regexPalabraReservada:
		return PALABRA_RESERVADA
	case regexOperadores:
		return OPERADOR_ARITMETICO // O OPERADOR_LOGICO, depende de la lógica de tu lenguaje
	case regexCadena:
		return CADENA_CARACTERES
	case regexComentario:
		return COMENTARIO_LINEA // O COMENTARIO_BLOQUE, depende de la lógica de tu lenguaje
	case regexOperadorAsignacion:
		return OPERADOR_ASIGNACION
	case regexParentesisApertura:
		return PARENTESIS_APERTURA
	case regexParentesisCierre:
		return PARENTESIS_CIERRE
	case regexLlaveApertura:
		return LLAVE_APERTURA
	case regexLlaveCierre:
		return LLAVE_CIERRE
	case regexTerminal:
		return TERMINAL
	case regexSeparador:
		return SEPARADOR
	case regexHexadecimal:
		return HEXADECIMAL
	default:
		return NO_RECONOCIDO
	}
}

func (al *AnalizadorLexico) AnalizarConExpresionesRegulares() {
	// Expresiones regulares
	regexIdentificador := regexp.MustCompile(`^[a-zA-Z_]\w{0,9}$`)
	regexNumeros := regexp.MustCompile(`^\d+(\.\d+)?$`)
	// ... Otras expresiones regulares ...

	i := 0

	for i < len(al.CodigoFuente) {
		// Intenta extraer el identificador
		token := al.extraerTokenConExpresionRegular(regexIdentificador, i)
		if token != nil {
			al.ListaTokens = append(al.ListaTokens, token)
			i = token.IndiceSiguiente
		} else {
			// Intenta extraer números
			tokenNumeros := al.extraerTokenConExpresionRegular(regexNumeros, i)
			if tokenNumeros != nil {
				al.ListaTokens = append(al.ListaTokens, tokenNumeros)
				i = tokenNumeros.IndiceSiguiente
			} else {
				// Intenta otras expresiones regulares ...
				// ... Otras expresiones regulares ...

				// Manejar otros casos...
				i++
			}
		}
	}

}

func (al *AnalizadorLexico) Analizar() {
	// Expresiones regulares
	regexIdentificador := regexp.MustCompile(`^[a-zA-Z_]\w{0,9}$`)
	//	regexNumeros := regexp.MustCompile(`^\d+(\.\d+)?$`)
	// ... Otras expresiones regulares ...

	i := 0

	for i < len(al.CodigoFuente) {
		// Intenta extraer el identificador usando expresiones regulares
		tokenExpReg := al.extraerTokenConExpresionRegular(regexIdentificador, i)
		if tokenExpReg != nil {
			al.ListaTokens = append(al.ListaTokens, tokenExpReg)
			i = tokenExpReg.IndiceSiguiente
		} else {
			// Si no se encontró un identificador con expresiones regulares, usa el AFD
			tokenAFD := al.ExtraerSgteToken(i)
			if tokenAFD != nil {
				al.ListaTokens = append(al.ListaTokens, tokenAFD)
				i = tokenAFD.IndiceSiguiente
			} else {
				// Intenta otras expresiones regulares ...
				// ... Otras expresiones regulares ...

				// Manejar otros casos...
				i++
			}
		}
	}
}

// ExtraerSgteToken extrae el próximo token del código fuente a partir de la posición dada
func (al *AnalizadorLexico) ExtraerSgteToken(indice int) *Token {
	var token *Token

	// Intenta extraer un entero
	token = al.extraerEntero(indice)
	if token != nil {
		return token
	}

	// Intenta otros métodos de extracción
	if token = al.extraerOperadorAritmetico(indice); token != nil {
		return token
	}
	if token = al.extraerOperadorRelacional(indice); token != nil {
		return token
	}
	if token = al.extraerOperadorLogico(indice); token != nil {
		return token
	}
	if token = al.extraerOperadorIncremento(indice); token != nil {
		return token
	}
	if token = al.extraerPalabraReservada(indice); token != nil {
		return token
	}
	if token = al.extraerCadenaCaracteres(indice); token != nil {
		return token
	}
	if token = al.extraerIdentificador(indice); token != nil {
		return token
	}

	// Extrae un token no reconocido si no se ha encontrado ningún tipo reconocido
	return al.extraerNoReconocido(indice)
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
func (al *AnalizadorLexico) GetListaTokens() []*Token {
	return al.ListaTokens
}

// Agrega a Categoria los operadores aritméticos
func (al *AnalizadorLexico) extraerOperadorAritmetico(indice int) *Token {
	operador := al.CodigoFuente[indice]
	if operador == '+' || operador == '-' || operador == '*' || operador == '/' {
		return &Token{
			Palabra:         string(operador),
			Categoria:       OPERADOR_ARITMETICO,
			IndiceSiguiente: indice + 1,
		}
	}
	return nil
}

// Agrega a Categoria los operadores relacionales
func (al *AnalizadorLexico) extraerOperadorRelacional(indice int) *Token {
	if indice+1 < len(al.CodigoFuente) {
		posibleOperador := al.CodigoFuente[indice : indice+2]
		if posibleOperador == "==" || posibleOperador == "!=" ||
			posibleOperador == "<=" || posibleOperador == ">=" ||
			posibleOperador == "<" || posibleOperador == ">" {
			return &Token{
				Palabra:         posibleOperador,
				Categoria:       OPERADOR_RELACIONAL,
				IndiceSiguiente: indice + 2,
			}
		}
	}
	return nil
}

// Agrega a Categoria los operadores lógicos
func (al *AnalizadorLexico) extraerOperadorLogico(indice int) *Token {
	operador := al.CodigoFuente[indice]
	if operador == '&' || operador == '|' {
		return &Token{
			Palabra:         string(operador),
			Categoria:       OPERADOR_LOGICO,
			IndiceSiguiente: indice + 1,
		}
	}
	return nil
}

// Agrega a Categoria los operadores de incremento/decremento
func (al *AnalizadorLexico) extraerOperadorIncremento(indice int) *Token {
	operador := al.CodigoFuente[indice]
	if operador == '+' || operador == '-' {
		return &Token{
			Palabra:         string(operador) + string(operador),
			Categoria:       OPERADOR_INCREMENTO,
			IndiceSiguiente: indice + 2,
		}
	}
	return nil
}

// Agrega a Categoria las palabras reservadas
func (al *AnalizadorLexico) extraerPalabraReservada(indice int) *Token {
	// Supongamos que tienes una lista de palabras reservadas
	palabrasReservadas := []string{"if", "else", "while", "int", "float", "boolean", "return"}
	for _, palabra := range palabrasReservadas {
		if strings.HasPrefix(al.CodigoFuente[indice:], palabra) && (indice+len(palabra) >= len(al.CodigoFuente) || !unicode.IsLetter(rune(al.CodigoFuente[indice+len(palabra)])) && !unicode.IsDigit(rune(al.CodigoFuente[indice+len(palabra)]))) {
			return &Token{
				Palabra:         palabra,
				Categoria:       PALABRA_RESERVADA,
				IndiceSiguiente: indice + len(palabra),
			}
		}
	}
	return nil
}

// Agrega a Categoria las cadenas de caracteres
func (al *AnalizadorLexico) extraerCadenaCaracteres(indice int) *Token {
	if al.CodigoFuente[indice] == '"' {
		posicion := indice + 1
		for indice < len(al.CodigoFuente) && al.CodigoFuente[indice] != '"' {
			indice++
		}
		if indice < len(al.CodigoFuente) && al.CodigoFuente[indice] == '"' {
			lexema := al.CodigoFuente[posicion:indice]
			return &Token{Palabra: lexema, Categoria: CADENA_CARACTERES, IndiceSiguiente: indice + 1}
		}
	}
	return nil
}

// Agrega a Categoria el identificador
func (al *AnalizadorLexico) extraerIdentificador(indice int) *Token {
	if unicode.IsLetter(rune(al.CodigoFuente[indice])) {
		posicion := indice
		for indice < len(al.CodigoFuente) &&
			(unicode.IsLetter(rune(al.CodigoFuente[indice])) || unicode.IsDigit(rune(al.CodigoFuente[indice])) || al.CodigoFuente[indice] == '_') {
			indice++
		}
		lexema := al.CodigoFuente[posicion:indice]
		return &Token{Palabra: lexema, Categoria: IDENTIFICADOR, IndiceSiguiente: indice}
	}
	return nil
}

// Agrega a Categoria los operadores de asignación
func (al *AnalizadorLexico) extraerOperadorAsignacion(indice int) *Token {
	if al.CodigoFuente[indice] == '=' {
		return &Token{
			Palabra:         "=",
			Categoria:       OPERADOR_ASIGNACION,
			IndiceSiguiente: indice + 1,
		}
	}
	return nil
}

// Agrega a Categoria los paréntesis de apertura
func (al *AnalizadorLexico) extraerParentesisApertura(indice int) *Token {
	if al.CodigoFuente[indice] == '(' {
		return &Token{
			Palabra:         "(",
			Categoria:       PARENTESIS_APERTURA,
			IndiceSiguiente: indice + 1,
		}
	}
	return nil
}

// Agrega a Categoria los paréntesis de cierre
func (al *AnalizadorLexico) extraerParentesisCierre(indice int) *Token {
	if al.CodigoFuente[indice] == ')' {
		return &Token{
			Palabra:         ")",
			Categoria:       PARENTESIS_CIERRE,
			IndiceSiguiente: indice + 1,
		}
	}
	return nil
}

// Agrega a Categoria las llaves de apertura
func (al *AnalizadorLexico) extraerLlaveApertura(indice int) *Token {
	if al.CodigoFuente[indice] == '{' {
		return &Token{
			Palabra:         "{",
			Categoria:       LLAVE_APERTURA,
			IndiceSiguiente: indice + 1,
		}
	}
	return nil
}

// Agrega a Categoria las llaves de cierre
func (al *AnalizadorLexico) extraerLlaveCierre(indice int) *Token {
	if al.CodigoFuente[indice] == '}' {
		return &Token{
			Palabra:         "}",
			Categoria:       LLAVE_CIERRE,
			IndiceSiguiente: indice + 1,
		}
	}
	return nil
}

// Agrega a Categoria el terminal
func (al *AnalizadorLexico) extraerTerminal(indice int) *Token {
	if al.CodigoFuente[indice] == ';' {
		return &Token{
			Palabra:         ";",
			Categoria:       TERMINAL,
			IndiceSiguiente: indice + 1,
		}
	}
	return nil
}

// Agrega a Categoria el separador
func (al *AnalizadorLexico) extraerSeparador(indice int) *Token {
	if al.CodigoFuente[indice] == ',' {
		return &Token{
			Palabra:         ",",
			Categoria:       SEPARADOR,
			IndiceSiguiente: indice + 1,
		}
	}
	return nil
}

// Agrega a Categoria el número hexadecimal
func (al *AnalizadorLexico) extraerHexadecimal(indice int) *Token {
	if strings.HasPrefix(al.CodigoFuente[indice:], "0x") {
		posicion := indice + 2
		indice += 2
		for indice < len(al.CodigoFuente) &&
			(unicode.IsDigit(rune(al.CodigoFuente[indice])) ||
				('a' <= al.CodigoFuente[indice] && al.CodigoFuente[indice] <= 'f') ||
				('A' <= al.CodigoFuente[indice] && al.CodigoFuente[indice] <= 'F')) {
			indice++
		}
		lexema := al.CodigoFuente[posicion:indice]
		return &Token{Palabra: lexema, Categoria: HEXADECIMAL, IndiceSiguiente: indice}
	}
	return nil
}

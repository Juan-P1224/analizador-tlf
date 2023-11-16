package modelo

// Token representa la estructura más pequeña de información de un analizador léxico (token)
type Token struct {
	Palabra         string
	Categoria       Categoria
	IndiceSiguiente int
}

// NewToken crea una nueva instancia de Token
func NewToken(palabra string, categoria Categoria, indiceSiguiente int) *Token {
	return &Token{
		Palabra:         palabra,
		Categoria:       categoria,
		IndiceSiguiente: indiceSiguiente,
	}
}

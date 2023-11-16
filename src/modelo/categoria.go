package modelo

type Categoria int

const (
	NO_RECONOCIDO Categoria = iota
	ENTERO
	DECIMAL
	IDENTIFICADOR
	PALABRA_RESERVADA
	CADENA_CARACTERES
	COMENTARIO_LINEA
	COMENTARIO_BLOQUE
	OPERADOR_ARITMETICO
	OPERADOR_RELACIONAL
	OPERADOR_LOGICO
	OPERADOR_INCREMENTO
)

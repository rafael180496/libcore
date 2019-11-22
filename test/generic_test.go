package test

import (
	"testing"

	utl "gitlab.com/gpctda/libcore/utility"
)

/*TestTrim : Quita los espacio de un texto */
func TestTrim(t *testing.T) {
	text := "Hola Mundo TDA"
	t.Logf("text:[%s]", text)
	t.Logf("trim:[%s]", utl.Trim(text))
}

/*TestUUID : Genera una clave unica*/
func TestUUID(t *testing.T) {
	text := utl.GeneredUUID()
	t.Logf("text:[%s]", text)
}

/*TestFloat64 : Redondea un valor float a x decimales*/
func TestFloat64(t *testing.T) {
	valor := 12.34661
	t.Logf("valor:[%f]", valor)
	valor = utl.RoundFloat64(valor, 2)
	t.Logf("valor:[%f]", valor)
}

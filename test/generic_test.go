package test

import (
	"fmt"
	"testing"

	utl "github.com/rafael180496/libcore/utility"
)

/*SubString : recorta una cadena */
func TestSubString(t *testing.T) {
	fmt.Printf("%s", utl.SubString("hola", 0, 2))
}

/*TestReturnIf : retorna con un condicional */
func TestReturnIf(t *testing.T) {
	t.Logf("%s", utl.ReturnIf(5 > 4, "It's true", "It's false"))
}

/*TestTrim : Quita los espacio de un texto */
func TestTrim(t *testing.T) {
	text := "Hola Mundo TDA"
	t.Logf("text:[%s]", text)
	t.Logf("trim:[%s]", utl.Trim(text))
}

/*TestFloat64 : Redondea un valor float a x decimales*/
func TestFloat64(t *testing.T) {
	valor := 12.34661
	t.Logf("valor:[%f]", valor)
	valor = utl.RoundFloat64(valor, 2)
	t.Logf("valor:[%f]", valor)
}

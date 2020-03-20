package utility

import (
	"fmt"
	ramdom "math/rand"
	"strings"
	"time"
	"unicode"

	uuid "github.com/satori/go.uuid"
)

/*ReturnIf : retorna un if ternario
https://github.com/TheMushrr00m/go-ternary Doc
ReturnIf(<bool expression>, <result for true>, <result for false>)
ReturnIf(5 > 4, "It's true", "It's false :(")
*/
func ReturnIf(a bool, b, c interface{}) interface{} {
	if a {
		return b
	}

	return c
}

/*CharRand : Genera una letra aleatoria.*/
func CharRand() string {
	ramdom.Seed(time.Now().UnixNano())
	return string(ramdom.Intn(122) + 97)
}

/*SubString : substring para un string en golang con runas*/
func SubString(cadena string, ini, cant int) string {
	runes := []rune(cadena)
	return string(runes[0:2])
}

/*Trim : Elimina el espacio de un string a nivel de runas.*/
func Trim(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {

			return -1
		}
		return r
	}, str)
}

/*IsSpace : valida si la cadena contiene espacio. */
func IsSpace(str string) bool {
	rest := false
	strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			rest = true
			return -1
		}
		return r
	}, str)
	return rest
}

/*GeneredUUID : Genera un codigo uuid unico */
func GeneredUUID() string {
	var err error
	u1 := uuid.Must(uuid.NewV4(), err)
	uuid := fmt.Sprintf("%s", u1)

	return uuid
}

package utility

import (
	"fmt"
	ramdom "math/rand"
	"strings"
	"time"
	"unicode"

	uuid "github.com/satori/go.uuid"
)

/*CharRand : Genera una letra aleatoria.*/
func CharRand() string {
	ramdom.Seed(time.Now().UnixNano())
	return string(ramdom.Intn(122) + 97)
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

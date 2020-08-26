package utility

import (
	"fmt"
	"math/rand"
	ramdom "math/rand"
	"strings"
	"time"
	"unicode"

	uuid "github.com/satori/go.uuid"
)

/*ValidDuplidArrayStr : valida un arreglo de string si estan duplicados*/
func ValidDuplidArrayStr(strs []string) bool {
	for i, str := range strs {
		for j, straux := range strs {
			if i != j && straux == str {
				return false
			}
		}
	}
	return true
}
func accStr(str, acc string) string {
	return ReturnIf("u" == acc, strings.ToUpper(acc), strings.ToLower(str)).(string)
}
func accStrs(acc string, strs ...string) []string {
	var strsNew []string
	for _, str := range strs {
		strsNew = append(strsNew, accStr(str, acc))
	}
	return strsNew
}

/*UpperStrs : coloca en mayusculas un arreglo compleot de strings*/
func UpperStrs(strs ...string) []string {
	return accStrs("u", strs...)
}

/*LowerStrs : coloca en minusculas un arreglo compleot de strings*/
func LowerStrs(strs ...string) []string {
	return accStrs("l", strs...)
}

/*InStr : compara  varios string con uno especifico*/
func InStr(str string, strs ...string) bool {
	for _, item := range strs {
		if str == item {
			return true
		}
	}
	return false
}

/*FilterExcl : excluye los registro de un arreglo A  con el B  ejemplo
A[a,b,c,d]
B[a,b]
Result
[c,d]
*/
func FilterExcl(strs []string, excl []string) []string {
	var ret []string
	for _, item := range strs {
		ind := true
		for _, item2 := range excl {
			if item2 == item {
				ind = false
			}
		}
		if ind {
			ret = append(ret, item)
		}
	}
	return ret
}

/*FilterStr : filtra un arreglo de string mediante un metodo definido */
func FilterStr(strs []string, valid func(string) bool) (ret []string) {
	for _, s := range strs {
		if valid(s) {
			ret = append(ret, s)
		}
	}
	return
}

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

/*StrRand : genera una cadena de caracteres ramdon*/
func StrRand(cant int, Upper bool) string {
	if cant <= 0 {
		cant = 1
	}
	str := ""
	for i := 0; i < cant; i++ {
		str += CharRand(Upper)
	}
	return str
}

/*CharRand : Genera una letra aleatoria upper indica si queres mayusculas o miniscula.*/
func CharRand(Upper bool) string {
	return ReturnIf(Upper, string(byte(RandInt(65, 90))), string(byte(RandInt(97, 122)))).(string)
}

/*RandInt : envia un numero aleatorio*/
func RandInt(min, max int) int {
	ramdom.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
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
	return fmt.Sprintf("%s", uuid.Must(uuid.NewV4(), err))
}

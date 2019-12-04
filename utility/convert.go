package utility

import (
	"encoding/hex"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

/*TimetoStr : Convierte un time al string con el formato YYYYMMDD que esta en constante.*/
func TimetoStr(dia time.Time) string {
	var (
		str strings.Builder
	)
	fmt.Fprintf(&str, FORMFE, dia.Year(), dia.Month(), dia.Day())
	return str.String()
}

/*StrtoByte : Convierte un string a array byte */
func StrtoByte(str string) []byte {

	return []byte(str)
}

/*BytetoStrHex : Convierte un   array byte a string */
func BytetoStrHex(str []byte) string {
	strByte := hex.EncodeToString(str)
	return strByte
}

/*SliceStrToStr : convierte un array de strings a un string con separador */
func SliceStrToStr(strs []string, separador string) string {
	str := ""
	for i := 0; i < len(strs); i++ {
		if i != (len(strs) - 1) {
			str += strs[i] + separador
		} else {
			str += strs[i]
		}
	}
	return str
}

/*StrtoByteHex : Convierte un string a array byte */
func StrtoByteHex(str string) []byte {
	strByte, _ := hex.DecodeString(str)
	return strByte
}

/*BytetoStr : Convierte un   array byte a string */
func BytetoStr(str []byte) string {
	return string(str)
}

/*RoundFloat64 : redondea un  */
func RoundFloat64(valor float64, dec int) float64 {
	if dec <= 0 {
		dec = 1
	}
	format := fmt.Sprintf(".%df", dec)
	format = "%" + format
	valstr := fmt.Sprintf(format, valor)
	val, err := StrToFloat64(valstr)
	if err != nil {
		return 0.00
	}
	return val

}

/*StrToFloat64 : Convierte un valor string a float64.*/
func StrToFloat64(val string) (float64, error) {
	Float64, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, SendErrorCod("CO01")
	}
	return Float64, nil
}

/*Float64ToStr : Convierte un valor float64 a string.*/
func Float64ToStr(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}

/*StrToInt64 : Convierte un string a int64 */
func StrToInt64(val string) (int64, error) {
	int64, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, SendErrorCod("CO02")
	}
	return int64, nil
}

/*Int64ToStr : Convierte un int64 a string.*/
func Int64ToStr(val int64) string {
	return strconv.FormatInt(val, 10)
}

/*StrToInt32 : Convierte un string a int32 */
func StrToInt32(val string) (int32, error) {
	int64, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, SendErrorCod("CO03")
	}
	return int32(int64), nil
}

/*Int32ToStr : Convierte un int32 a string.*/
func Int32ToStr(val int32) string {
	return strconv.FormatInt(int64(val), 10)
}

/*IntToStr : Convierte un int a string.*/
func IntToStr(val int) string {
	return strconv.FormatInt(int64(val), 10)
}

/*MapKeys : Conviert un map a interface o datos dinamicos.*/
func MapKeys(data *map[interface{}]interface{}) []interface{} {

	keys := make([]interface{}, len(*data))
	i := 0
	for key := range *data {
		keys[i] = key
		i++
	}
	return keys
}

/*MapStrKeys : Convierte un map a arreglo de string*/
func MapStrKeys(data *map[string]interface{}) []string {

	keys := make([]string, len(*data))
	i := 0
	for key := range *data {
		keys[i] = key
		i++
	}
	return keys

}

/*StructToMap : Convierte un struct a map[string] */
func StructToMap(i interface{}) (valores url.Values) {
	valores = url.Values{}
	atributo := reflect.ValueOf(i).Elem()
	tipo := atributo.Type()
	for i := 0; i < atributo.NumField(); i++ {
		f := atributo.Field(i)
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		valores.Set(tipo.Field(i).Name, v)
	}
	return
}

package database

import (
	"encoding/json"
	"fmt"
	"strconv"

	utl "github.com/rafael180496/libcore/utility"
)

type (

	/*StSQLData : Estructura para manipulacion de columnas JSON esto funciona para capturar una sola respuesta a nivel de solucion
	ejemplo
	select getdata() as data --> donde data deve ser un json o []byte
	*/
	StSQLData struct {
		Data []byte `db:"data"`
	}

	/*StQuery : Estructura para ejecutar query de base de datos. */
	StQuery struct {
		Querie string `json:"querie"`
		Args   map[string]interface{}
	}

	/*StData : Estructura que extrae los datos de una consulta de base de datos tramformandola en map*/
	StData map[string]interface{}
)

/*ToJSON : Convierte un  StSQLData a utl.JSON*/
func (p *StSQLData) ToJSON() utl.JSON {
	var data utl.JSON
	data = p.Data
	return data
}

/*ToMap : Convierte un StSQLData a map interface*/
func (p *StSQLData) ToMap() ([]map[string]interface{}, error) {
	return utl.JSONtoObj(p.ToJSON())
}

/*Unmarshal : captura una structura*/
func (p *StSQLData) Unmarshal(v interface{}) error {
	return json.Unmarshal(p.Data, v)
}

/*Filter : Excluye key de un map interface*/
func (p *StData) Filter(keys ...string) StData {
	datanew := make(StData)
	clone := *p
	for k, vl := range clone {
		if !utl.InStr(k, keys...) {
			datanew[k] = vl
		}
	}
	return datanew
}

/*ValidColum : valida si un campo existe*/
func (p *StData) ValidColum(col string) bool {
	for _, item := range p.KeyColum() {
		if item == col {
			return true
		}
	}
	return false
}

/*KeyColum : envia las columnas que contiene la data*/
func (p *StData) KeyColum() []string {
	var colums []string
	clone := *p
	for k := range clone {
		colums = append(colums, k)
	}
	return colums
}

/*ToJSON : Convierte la estructura  StData en JSON para envios externos a rest api.*/
func (p *StData) ToJSON() ([]byte, error) {
	jsonData, err := json.Marshal(&p)
	if err != nil {
		return nil, utl.Msj.GetError("CN01")
	}
	return jsonData, nil
}

/*ToString : Convierte el valor del map interface{} a string.*/
func (p *StData) ToString(columna string) (string, error) {
	var valor interface{}
	valor = (*p)[columna]

	switch valor.(type) {
	case string:
		return valor.(string), nil
	case int:
		return strconv.Itoa(valor.(int)), nil
	case int32:
		return utl.Int32ToStr(valor.(int32)), nil
	case int64:
		return utl.Int64ToStr(valor.(int64)), nil
	case float64:
		return utl.Float64ToStr(valor.(float64)), nil
	case float32:
		return utl.Float64ToStr(float64(valor.(float32))), nil

	}
	return fmt.Sprintf("%v", valor), utl.Msj.GetError("CN02")
}

/*ToInt : Convierte el valor del map interface{} a int.*/
func (p *StData) ToInt(columna string) (int, error) {
	var valor interface{}
	valor = (*p)[columna]

	switch valor.(type) {
	case string:
		return strconv.Atoi(valor.(string))
	case int:
		return valor.(int), nil
	case int32:
		return int(valor.(int32)), nil
	case int64:
		return int(valor.(int64)), nil

	}

	return 0, utl.Msj.GetError("CN03")
}

/*ToBool : Convierte el valor del map interface{} a bool.*/
func (p *StData) ToBool(columna string) bool {
	var valor interface{}
	valor = (*p)[columna]
	switch valor.(type) {
	case bool:
		return valor.(bool)
	default:
		return false
	}
}

/*ToInt64 : Convierte el valor del map interface{} a int64.*/
func (p *StData) ToInt64(columna string) (int64, error) {
	var valor interface{}
	valor = (*p)[columna]

	switch valor.(type) {
	case string:
		return utl.StrToInt64(valor.(string))
	case int:
		return int64(valor.(int)), nil
	case int32:
		return int64(valor.(int32)), nil
	case int64:
		return valor.(int64), nil
	}
	return 0, utl.Msj.GetError("CN03")
}

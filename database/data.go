package database

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	utl "github.com/rafael180496/libcore/utility"

	"github.com/jmoiron/sqlx"
)

type (
	/*StQuery : Estructura para ejecutar query de base de datos. */
	StQuery struct {
		Querie string `json:"querie"`
		Args   map[string]interface{}
	}

	/*StData : Estructura que extrae los datos de una consulta de base de datos tramformandola en map*/
	StData map[string]interface{}
)

/*ToJSON : Convierte la estructura  StData en JSON para envios externos a rest api.*/
func (p *StData) ToJSON() ([]byte, error) {
	jsonData, err := json.Marshal(&p)
	if err != nil {
		return nil, utl.SendErrorCod("CN01")
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
	return fmt.Sprintf("%v", valor), utl.SendErrorCod("CN02")
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

	return 0, utl.SendErrorCod("CN03")
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
	return 0, utl.SendErrorCod("CN03")
}

/*envTableData : convierte los datos en tipos  json */
func envTableData(tableData []map[string]interface{}, datachan chan utl.JSON, errchan chan error) {
	d, err := utl.NewJSON(tableData)
	if err != nil {
		errchan <- err
	} else {
		datachan <- d
	}
}

/*ValidSelect : validaciones basicas de un select  */
func ValidSelect(sql string) error {
	sql = strings.ToLower(sql)
	if strings.Replace(sql, " ", "", -1) == "" {
		return utl.SendErrorCod("CN05")
	}
	if strings.Contains(sql, INSERT) || strings.Contains(sql, DELETE) || strings.Contains(sql, UPDATE) {
		return utl.SendErrorCod("CN05")
	}
	if !strings.Contains(sql, SELECT) || !strings.Contains(sql, FROM) {
		return utl.SendErrorCod("CN05")
	}

	return nil
}

/*validTipDB : valida el tipo de sql que insertara.*/
func validTipDB(sqlOrig string, tipo string) error {
	var (
		valid bool
	)
	sqlOrig = strings.ToLower(sqlOrig)

	switch tipo {
	case INSERT:
		valid = strings.Contains(sqlOrig, INSERT)
		if !valid {
			return utl.SendErrorCod("CN04")
		}
		return nil
	case UPDATE, DELETE:
		valid = strings.Contains(sqlOrig, UPDATE)
		if !valid {
			valid = strings.Contains(sqlOrig, DELETE)
			if !valid {
				return utl.SendErrorCod("CN05")
			}
		}
		return nil
	default:
		return utl.SendErrorCod("CN05")
	}

}

/*sendData : captura los datos de la tabla*/
func sendData(valores []interface{}, columnas []string) StData {
	data := make(StData)
	for i, col := range valores {
		if col == nil {
			continue
		}
		switch col.(type) {

		case []byte:
			data[columnas[i]] = string(col.([]byte))
		default:
			data[columnas[i]] = col
		}
	}
	return data
}

/*scanData : escanea las fila regresando un tipo generico */
func scanData(rows *sqlx.Rows, cantrow int) ([]StData, error) {
	var (
		result   []StData
		columnas []string
		err      error
		controws = 0
	)
	columnas, err = rows.Columns()
	if err != nil {
		return result, utl.SendErrorCod("CN06")
	}
	ptrData := make([]interface{}, len(columnas))
	valores := make([]interface{}, len(columnas))
	for i := range valores {
		ptrData[i] = &valores[i]
	}
	for rows.Next() {
		if cantrow > 0 {
			if controws > cantrow || cantrow == 0 {
				break
			}
			controws++
		}
		err = rows.Scan(ptrData...)
		if err != nil {
			return result, utl.SendErrorCod("CN07")
		}
		data := sendData(valores, columnas)
		result = append(result, data)
	}
	return result, nil
}

/*ScanDataGeneric : escanea los datos de ejecutarquerieGO en rutinas */
func ScanDataGeneric(filas *sqlx.Rows, columnas []string, cantrows int, datachan chan utl.JSON, errchan chan error) {
	defer filas.Close()
	tableData := []map[string]interface{}{}
	values := make([]interface{}, len(columnas))
	valuePtrs := make([]interface{}, len(columnas))
	for i := 0; i < len(columnas); i++ {
		valuePtrs[i] = &values[i]
	}
	for filas.Next() {
		err := filas.Scan(valuePtrs...)
		if err != nil {
			errchan <- err
			break
		}
		entrada := make(map[string]interface{})
		for i, col := range columnas {
			var v interface{}
			val := values[i]
			switch vv := val.(type) {
			case []byte:
				v = string(vv)
			default:
				v = vv
			}
			entrada[col] = v
		}
		tableData = append(tableData, entrada)
		if cantrows > 0 && len(tableData) >= cantrows {
			envTableData(tableData, datachan, errchan)
			tableData = []map[string]interface{}{}
			break
		}
	}
	if filas.Err() != nil {
		errchan <- filas.Err()
	}
	close(datachan)
	close(errchan)
}

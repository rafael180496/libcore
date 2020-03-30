package database

import (
	"strings"

	"github.com/jmoiron/sqlx"
	utl "github.com/rafael180496/libcore/utility"
)

/*envTableData : convierte los datos en tipos  json */
func envTableData(tableData []map[string]interface{}, datachan chan utl.JSON, errchan chan error) {
	d, err := utl.NewJSON(tableData)
	if err != nil {
		errchan <- err
	} else {
		datachan <- d
	}
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
			return utl.Msj.GetError("CN04")
		}
		return nil
	case UPDATE, DELETE:
		valid = strings.Contains(sqlOrig, UPDATE)
		if !valid {
			valid = strings.Contains(sqlOrig, DELETE)
			if !valid {
				return utl.Msj.GetError("CN05")
			}
		}
		return nil
	default:
		return utl.Msj.GetError("CN05")
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
		return result, utl.Msj.GetError("CN06")
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
			return result, utl.Msj.GetError("CN07")
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

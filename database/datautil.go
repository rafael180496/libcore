package database

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	utl "github.com/rafael180496/libcore/utility"
	"gopkg.in/ini.v1"
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

/*strURL : Arma la cadena de conexion dependiendo del tipo*/
func strURL(tipo string, conexion StCadConect) (string, string) {
	switch tipo {
	case Ora:
		/*Open("oracle", "oracle://user:pass@server/service_name")*/
		return PrefijosDB[tipo], fmt.Sprintf(CADCONN[tipo], conexion.Usuario, conexion.Clave, conexion.Host, conexion.Puerto, conexion.Nombre)
	case Post:
		/*"postgres://user:password@localhost:port/bdnamme?sslmode=verify-full"*/
		return PrefijosDB[tipo], fmt.Sprintf(CADCONN[tipo], conexion.Usuario, conexion.Clave, conexion.Host, conexion.Puerto, conexion.Nombre, conexion.Sslmode)
	case Mysql:
		/*sql.Open("mssql", "user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true") */
		return PrefijosDB[tipo], fmt.Sprintf(CADCONN[tipo], conexion.Usuario, conexion.Clave, conexion.Host, conexion.Puerto, conexion.Nombre)
	case Sqlser:
		return PrefijosDB[tipo], fmt.Sprintf(CADCONN[tipo], conexion.Host, conexion.Usuario, conexion.Clave, conexion.Puerto, conexion.Nombre)
	case SQLLite:
		return PrefijosDB[tipo], fmt.Sprintf(CADCONN[tipo], conexion.File)
	default:
		return "", ""
	}
}

/*validTipDB : valida el tipo de sql que insertara.*/
func validTipDB(sqlOrig string, tipo string) error {
	var (
		valid bool
	)
	sqlOrig = strings.ToUpper(sqlOrig)

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
func scanData(rows *sqlx.Rows, maxRows int, indLimit bool) ([]StData, error) {
	var (
		result    []StData
		columnas  []string
		err       error
		countRows = 0
	)
	maxRows = utl.ReturnIf(maxRows <= 0, 1, maxRows).(int)
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
		if indLimit == true {
			if countRows > maxRows {
				break
			}
			countRows++
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

/*validTp : valida los tipos de conexion disponible*/
func validTp(tp string) bool {
	switch tp {
	case SQLLite, Ora, Post, Mysql, Sqlser:
		return true
	default:
		return false
	}
}
func readIni(source interface{}) (StCadConect, error) {
	var cnx StCadConect
	cfg, err := ini.Load(source)
	if err != nil {
		return cnx, utl.Msj.GetError("CN11")
	}
	err = cfg.Section("database").MapTo(&cnx)
	if err != nil {
		return cnx, err
	}
	if !cnx.ValidCad() {
		return cnx, utl.Msj.GetError("CN12")
	}
	return cnx, nil
}

/*DecripConect : desencripta una conexion de base de datos .ini con una encriptacion AES256 creada del mismo
paquete utility*/
func DecripConect(data []byte, pass string) (StCadConect, error) {
	var cnx StCadConect
	dataNew, err := utl.DesencripAES(pass, utl.BytetoStr(data))
	if err != nil {
		return cnx, err
	}
	cnx, err = readIni([]byte(dataNew))
	if err != nil {
		return cnx, err
	}
	return cnx, nil
}

/*CreateDBConect : Crea una conexion de base de datos valida y la genera como un .db con una clave aes*/
func CreateDBConect(cnx StCadConect, pass string) ([]byte, error) {
	pass = utl.Trim(pass)
	if !utl.IsNilStr(pass) {
		return nil, utl.StrErr("La clave esta vacia por favor introducir una clave")
	}
	if !cnx.ValidCad() {
		return nil, utl.StrErr("La conexion no pasa las validaciones.")
	}
	cfg := ini.Empty()
	sec, err := cfg.NewSection("database")
	if err != nil {
		return nil, err
	}
	err = sec.ReflectFrom(&cnx)
	if err != nil {
		return nil, err
	}
	cfg.DeleteSection("DEFAULT")
	var buf bytes.Buffer
	cfg.WriteTo(&buf)
	data := buf.String()
	dataencrip, err := utl.EncripAES(pass, data)
	if err != nil {
		return nil, err
	}
	return utl.StrtoByte(dataencrip), nil
}

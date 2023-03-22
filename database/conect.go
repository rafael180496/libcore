package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	utl "github.com/rafael180496/core-util/utility"

	/*Conexion a mysql*/
	_ "github.com/go-sql-driver/mysql"
	/*Conexion a postgrest*/
	_ "github.com/lib/pq"
	/*Conexion a oracle*/
	_ "gopkg.in/rana/ora.v4"
	/*Conexion a sql server*/
	_ "github.com/denisenkom/go-mssqldb"
	/*Conexion a sqllite*/
	_ "github.com/mattn/go-sqlite3"
)

type (
	/*StCadConect : Estructura para generar la cadena de  conexiones de base de datos */
	StCadConect struct {
		File    string `json:"filedb"  ini:"filedb"`
		Usuario string `json:"usuario" ini:"usuario"`
		Clave   string `json:"clave"   ini:"clave"`
		Nombre  string `json:"nombre"  ini:"nombre"`
		Tipo    string `json:"tipo"    ini:"tipo"`
		Host    string `json:"host"    ini:"host"`
		Puerto  int    `json:"puerto"  ini:"puerto"`
		Sslmode string `json:"sslmode" ini:"sslmode"`
	}
	/*StConect : Estructura que contiene la conexion a x tipo de base de datos.*/
	StConect struct {
		Conexion     StCadConect
		urlNative    string
		DBGO         *sqlx.DB
		DBTx         *sql.Tx
		DBStmt       *sql.Stmt
		backupScript string
		Queries      map[string]string
	}
)

/*SetBackupScript : setea un scrip backup para la creacion de base de datos en modelos go*/
func (p *StConect) SetBackupScript(sql string) {
	p.backupScript = sql
}

/*ExecBackup : ejecuta el querie backup */
func (p *StConect) ExecBackup() error {
	if len(p.backupScript) <= 0 {
		return utl.Msj.GetError("CN22")
	}
	err := p.Con()
	if err != nil {
		return err
	}
	tx := p.DBGO.MustBegin()
	_, err = tx.Exec(p.backupScript)
	if err != nil {
		p.Close()
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		p.Close()
		tx.Rollback()
		return err
	}
	return nil
}

/*SendSQL : envia un sql con los argumentos */
func (p *StConect) SendSQL(code string, args map[string]interface{}) StQuery {
	return StQuery{
		Querie: p.Queries[code],
		Args:   args,
	}
}

/*Close : cierra las conexiones de base de datos intanciadas*/
func (p *StConect) Close() error {
	if p.DBGO == nil {
		return nil
	}
	err := p.DBGO.Close()
	if err != nil {
		return err
	}
	return nil
}

/*NamedIn : procesa los argumentos y sql para agarrar la clausula IN */
func (p *StConect) NamedIn(query StQuery) (string, []interface{}, error) {
	var (
		sqltemp string
		args    []interface{}
		err     error
	)
	sqltemp, args, err = sqlx.Named(query.Querie, query.Args)
	if err != nil {
		return "", nil, err
	}
	sqltemp, args, err = sqlx.In(sqltemp, args...)
	if err != nil {
		return "", nil, err
	}
	sqltemp = p.DBGO.Rebind(sqltemp)

	return sqltemp, args, err
}

/*Trim : Elimina los espacio en cualquier campo string */
func (p *StCadConect) Trim() {
	p.File = utl.Trim(p.File)
	p.Usuario = utl.Trim(p.Usuario)
	p.Clave = utl.Trim(p.Clave)
	p.Nombre = utl.Trim(p.Nombre)
	p.Tipo = utl.Trim(p.Tipo)
	p.Host = utl.Trim(p.Host)
	p.Sslmode = utl.Trim(p.Sslmode)
	if p.Tipo == Post && p.Sslmode == "" {
		p.Sslmode = Ssmodes[0]
	}

}

/*ConfigURL : captura una conexion nativa de drive para base de datos*/
func (p *StConect) ConfigURL(url string) {
	p.urlNative = url
}

/*ConfigJSON : Lee las configuraciones de conexion mediante un .json

Ejemplo:

{

	"usuario":"prueba",
	"clave":"prueba",
	"nombre":"prueba",
	"tipo":"POST",
	"host":"Localhost",
	"puerto":3000,
	"sslmode":"",
	"filedb":""

}

*/
func (p *StConect) ConfigJSON(PathJSON string) error {
	var (
		err     error
		cad     StCadConect
		ptrArch *os.File
	)
	if !utl.FileExt(PathJSON, "JSON") {
		return utl.Msj.GetError("CN09")
	}
	PathJSON, err = utl.TrimFile(PathJSON)
	if err != nil {
		return utl.Msj.GetError("CN08")
	}
	ptrArch, err = os.Open(PathJSON)
	if err != nil {
		return utl.Msj.GetError("CN08")
	}
	defer ptrArch.Close()
	decJSON := json.NewDecoder(ptrArch)
	err = decJSON.Decode(&cad)
	if err != nil {
		return utl.Msj.GetError("CN08")
	}
	if !cad.ValidCad() {
		return utl.Msj.GetError("CN12")
	}
	p.Conexion = cad
	return nil
}

/*ConfigDBX : Lee las configuraciones de conexion mediante un archivo encriptado .dbx este se debe enviar la clave*/
func (p *StConect) ConfigDBX(path, pass string) error {
	if !utl.FileExt(path, "DBX") {
		return utl.StrErr("No existe el archivo .dbx")
	}
	dataraw, err := utl.ReadFileStr(path)
	if err != nil {
		return err
	}
	cad, err := DecripConect(utl.StrtoByte(dataraw), pass)
	if err != nil {
		return err
	}
	p.Conexion = cad
	return nil
}

/*ConfigINI : Lee las configuraciones de conexion mediante un .ini

Ejemplo:

[database]

usuario = prueba

clave = prueba

nombre  = prueba

tipo = POST

puerto = 5433

host = Localhost

sslmode = opcional

filedb = opcional sqllite

*/
func (p *StConect) ConfigINI(PathINI string) error {
	if !utl.FileExt(PathINI, "INI") {
		return utl.Msj.GetError("CN10")
	}
	cad, err := readIni(PathINI)
	if err != nil {
		return err
	}
	p.Conexion = cad
	return nil
}

/*ConfigENV : lee las configuracion de la base de datos mediante variables de entorno
Ejemplo:
ENV USUARIO = prueba
ENV CLAVE = prueba
ENV NOMBRE  = prueba
ENV TIPO = POST
ENV PUERTO = 5433
ENV HOST = Localhost
ENV SSLMODE = opcional
ENV  FILEDB = opcional sqllite
*/
func (p *StConect) ConfigENV() error {
	var (
		cad StCadConect
	)
	cad.Clave = os.Getenv("CLAVE")
	cad.Usuario = os.Getenv("USUARIO")
	cad.Nombre = os.Getenv("NOMBRE")
	cad.Tipo = os.Getenv("TIPO")
	cad.Puerto = utl.ToInt(os.Getenv("TIPO"))
	cad.Host = os.Getenv("HOST")
	cad.Sslmode = os.Getenv("SSLMODE")
	cad.File = os.Getenv("FILEDB")
	if !cad.ValidCad() {
		return utl.Msj.GetError("CN12")
	}
	p.Conexion = cad
	return nil
}

/*ResetCnx : Limpia la cadena de conexion*/
func (p *StConect) ResetCnx() {
	p.Conexion = StCadConect{}
}

/*ToString : Muestra la estructura  StCadConect*/
func (p *StCadConect) ToString() string {
	return fmt.Sprintf(FORMATTOSTRCONECT, p.Clave, p.Host, p.Nombre, p.Puerto, p.Sslmode, p.Tipo, p.Usuario, p.File)
}

/*ValidCad : valida la cadena de conexion capturada */
func (p *StCadConect) ValidCad() bool {
	p.Trim()
	if !validTp(p.Tipo) {
		return false
	}
	if p.Tipo != SQLLite && (!utl.IsNilArrayStr(p.Clave, p.Usuario, p.Nombre, p.Tipo, p.Host) || p.Puerto <= 0) {
		return false
	}
	if p.Tipo == SQLLite && !utl.IsNilStr(p.File) {
		return false
	}
	return true
}

/*Con : Crear una conexion ala base de datos configurada en la cadena.*/
func (p *StConect) Con() error {
	var (
		err, errping error
	)
	conexion := p.Conexion
	prefijo, cadena := strURL(p.Conexion.Tipo, conexion)
	cadena = utl.ReturnIf(!utl.IsNilStr(p.urlNative), cadena, p.urlNative).(string)
	if cadena == "" {
		return utl.Msj.GetError("CN13")
	}
	if p.DBGO != nil {
		errping = p.DBGO.Ping()
	}
	if errping != nil || p.DBGO == nil {
		if p.Conexion.Tipo == SQLLite && p.createDB() != nil {
			return utl.Msj.GetError("CN20")
		}
		p.DBGO, err = sqlx.Connect(prefijo, cadena)
		if err != nil {
			return utl.Msj.GetError("CN14")
		}
	}
	return nil
}

/*Insert : Inserta a cualquier tabla donde esta conectado devuelve true si fue guardado o false si no guardo nada.*/
func (p *StConect) Insert(Data []StQuery) error {
	return p.ExecValid(Data, INSERT)
}

/*UpdateOrDelete : actualiza e elimina a cualquier tabla donde esta conectado devuelve la cantidad de filas afectadas.*/
func (p *StConect) UpdateOrDelete(Data []StQuery) (int64, error) {
	err := p.ExecValid(Data, DELETE)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

/*ExecDatatable : ejecuta a nivel de base de datos una accione datable esta puede ser INSERT,DELETE,UPDATE*/
func (p *StConect) ExecDatatable(data DataTable, acc string, indConect bool) error {
	queries, err := data.GenSQL(acc)
	if err != nil {
		return err
	}
	err = p.Exec(queries, indConect)
	if err != nil {
		return err
	}
	return nil
}

/*Exec :Ejecuta una accion de base de datos nativa con rollback*/
func (p *StConect) Exec(Data []StQuery, indConect bool) error {
	return p.execAux(Data, "", false, indConect)
}

/*ExecOne :Ejecuta un StQuery navito haciendo rollback con un error*/
func (p *StConect) ExecOne(Data StQuery, indConect bool) error {
	err := p.Con()
	if err != nil {
		return err
	}
	//Bloque de ejecucion
	tx := p.DBGO.MustBegin()
	_, err = tx.NamedExec(Data.Querie, Data.Args)
	if err != nil {
		p.Close()
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		p.Close()
		tx.Rollback()
		return err
	}
	if !indConect {
		p.Close()
	}
	return nil
}

/*ExecValid :Ejecuta una accion de base de datos nativa con rollback y validacion de insert e delete o que tipo de accion es */
func (p *StConect) ExecValid(Data []StQuery, tipacc string) error {
	return p.execAux(Data, tipacc, true, false)
}

/*ExecNative :  ejecuta la funcion nativa del paquete sql*/
func (p *StConect) ExecNative(sql string, indConect bool, args ...interface{}) (sql.Result, error) {
	if !utl.IsNilStr(sql) {
		return nil, utl.StrErr("El querie esta vacio")
	}
	err := p.Con()
	if err != nil {
		return nil, err
	}
	tx := p.DBGO.MustBegin()
	rel, err := tx.Exec(sql, args...)
	if err != nil {
		p.Close()
		tx.Rollback()
		return rel, err
	}
	err = tx.Commit()
	if err != nil {
		p.Close()
		tx.Rollback()
		return nil, err
	}
	if !indConect {
		p.Close()
	}
	return rel, nil
}

/*Test : Valida si se puede conectar ala base de datos antes de un  uso.*/
func (p *StConect) Test() bool {
	prueba := new(StQuery)
	switch p.Conexion.Tipo {
	case Post, Mysql, Sqlser, SQLLite:
		prueba.Querie = `SELECT 1`
	case Ora:
		prueba.Querie = `SELECT 1 FROM DUAL`
	}
	dato, err := p.QueryMap(*prueba, 1, false, true)
	if err != nil || len(dato) <= 0 {
		return false
	}
	return true
}

/*ValidTable : valida si la tabla a buscar existe*/
func (p *StConect) ValidTable(table string) bool {
	prueba := StQuery{
		Querie: TESTTABLE[p.Conexion.Tipo],
		Args: map[string]interface{}{
			"TABLENAME": table,
		},
	}
	dato, err := p.QueryMap(prueba, 1, false, true)
	if err != nil || len(dato) <= 0 {
		return false
	}
	num, erraux := dato[0].ToInt("REG")
	if num <= 0 || erraux != nil {
		return false
	}
	return true
}

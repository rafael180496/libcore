package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/go-ini/ini"
	"github.com/jmoiron/sqlx"
	utl "github.com/rafael180496/libcore/utility"

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
		File    string `json:"filedb"`
		Usuario string `json:"usuario"`
		Clave   string `json:"clave"`
		Nombre  string `json:"nombre"`
		Tipo    string `json:"tipo"`
		Host    string `json:"host"`
		Puerto  int    `json:"puerto"`
		Sslmode string `json:"sslmode"`
	}
	/*StConect : Estructura que contiene la conexion a x tipo de base de datos.*/
	StConect struct {
		Conexion StCadConect
		DBGO     *sqlx.DB
		DBTx     *sql.Tx
		DBStmt   *sql.Stmt
	}
)

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
	if !utl.FileExist(PathJSON, false) || !utl.FileExt(PathJSON, "JSON") {
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
	var (
		cad StCadConect
	)
	if !utl.FileExist(PathINI, false) || !utl.FileExt(PathINI, "INI") {
		return utl.Msj.GetError("CN10")
	}

	iniArch, err := ini.Load(PathINI)
	if err != nil {
		return utl.Msj.GetError("CN11")
	}
	cad.Usuario = iniArch.Section("database").Key("usuario").String()
	cad.Clave = iniArch.Section("database").Key("clave").String()
	cad.Nombre = iniArch.Section("database").Key("nombre").String()
	cad.Tipo = iniArch.Section("database").Key("tipo").String()
	cad.File = iniArch.Section("database").Key("filedb").String()
	cad.Puerto, err = iniArch.Section("database").Key("puerto").Int()
	if err != nil {
		cad.Puerto = 0
	}
	cad.Host = iniArch.Section("database").Key("host").String()
	/*opcional configuracion temporal*/
	cad.Sslmode = iniArch.Section("database").Key("sslmode").String()
	if !cad.ValidCad() {
		return utl.Msj.GetError("CN12")
	}
	p.Conexion = cad
	return nil
}

/*TranSQL : procesa los prefijos que le toca a cada tipo de base de datos.*/
func (p *StConect) TranSQL(SQL string) string {
	return strings.ToLower(strings.Replace(SQL, PrefixG, Prefijos[p.Conexion.Tipo], -1))
}

/*ToString : Muestra la estructura  StCadConect*/
func (p *StCadConect) ToString() string {
	return fmt.Sprintf(FORMATTOSTRCONECT, p.Clave, p.Host, p.Nombre, p.Puerto, p.Sslmode, p.Tipo, p.Usuario, p.File)
}

/*ValidCad : valida la cadena de conexion capturada */
func (p *StCadConect) ValidCad() bool {
	p.Trim()
	if p.Tipo != SQLLite && (p.Clave == "" || p.Usuario == "" || p.Nombre == "" || p.Tipo == "" || p.Host == "" || p.Puerto <= 0) {
		return false
	}
	if p.Tipo == SQLLite && p.File == "" {
		return false
	}
	return true
}

/*strURL : Arma la cadena de conexion dependiendo del tipo*/
func strURL(tipo string, conexion StCadConect) (string, string) {
	switch tipo {
	case Ora:
		/*Open("ora", "user/passw@host:port/sid")*/
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

/*Con : Crear una conexion ala base de datos configurada en la cadena.*/
func (p *StConect) Con() error {
	var (
		err, errping error
	)
	conexion := p.Conexion
	prefijo, cadena := strURL(p.Conexion.Tipo, conexion)
	if cadena == "" {
		return utl.Msj.GetError("CN13")
	}
	if p.DBGO != nil {
		errping = p.DBGO.Ping()
	}
	if errping != nil || p.DBGO == nil {
		if (p.Conexion.Tipo == SQLLite) && (!utl.FileExist(conexion.File, false)) {
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
	_, err := p.Exec(Data, INSERT)
	if err != nil {
		return err
	}
	return nil
}

/*UpdateOrDelete : actualiza e elimina a cualquier tabla donde esta conectado devuelve la cantidad de filas afectadas.*/
func (p *StConect) UpdateOrDelete(Data []StQuery) (int64, error) {
	rels, err := p.Exec(Data, DELETE)
	if err != nil {
		return 0, err
	}
	filasCont, _ := rels.RowsAffected()
	return filasCont, nil
}

/*Exec : Ejecuta una accion en la conexion de base de datos es la funcion base para las funciones InsertQuery  UpdateOrDeleteQuery. */
func (p *StConect) Exec(Data []StQuery, tipDB string) (sql.Result, error) {
	var (
		err    error
		result sql.Result
	)

	FinChan := make(chan bool)
	defer close(FinChan)
	go func() {
		err = p.Con()
		if err != nil {
			err = utl.Msj.GetError("CN14")
			FinChan <- true
			return
		}
		tx := p.DBGO.MustBegin()

		for _, dat := range Data {
			err = validTipDB(dat.Querie, tipDB)
			if err != nil {
				p.Close()
				tx.Rollback()
				FinChan <- true
				return
			}
			result, err = tx.NamedExec(dat.Querie, dat.Args)
			if err != nil {

				p.Close()
				tx.Rollback()
				FinChan <- true
				return
			}
		}
		err = tx.Commit()
		if err != nil {
			p.Close()
			tx.Rollback()
			err = utl.Msj.GetError("CN17")
			FinChan <- true
			return
		}
		p.Close()
		FinChan <- true
		return
	}()

	for {
		select {
		default:
			continue
		case <-FinChan:
			if err != nil {
				return result, err
			}

			return result, nil
		}
	}

}

/*QueryStruct : Ejecuta un query en la base de datos y
  captura la data con struct
	EjecutarQueryStruct(&data,sql,true)
	indConect = true deja la conexion abierta
*/
func (p *StConect) QueryStruct(datadest interface{}, query StQuery, indConect bool) error {
	var (
		err     error
		args    []interface{}
		sqltemp string
	)

	FinChan := make(chan bool)
	defer close(FinChan)
	go func() {

		err = p.Con()
		if err != nil {
			err = utl.Msj.GetError("CN14")
			FinChan <- true
			return
		}
		sqltemp, args, err = p.NamedIn(query)
		if err != nil {
			FinChan <- true
			return
		}
		err = p.DBGO.Select(datadest, sqltemp, args...)
		if err != nil {
			p.Close()
			FinChan <- true
			return
		}
		if !indConect {
			p.Close()
		}
		FinChan <- true
		return
	}()
	for {
		select {
		default:
			continue
		case <-FinChan:

			if err != nil {

				return err
			}

			return nil
		}
	}

}

/*QueryRows : Ejecuta un query en la base de datos y
  devuelve un puntero de *Rows de sqlx
	indConect = true deja la conexion abierta
*/
func (p *StConect) QueryRows(query StQuery, indConect bool) (*sqlx.Rows, error) {
	var (
		err     error
		filas   *sqlx.Rows
		sqltemp string
		args    []interface{}
	)

	FinChan := make(chan bool)
	defer close(FinChan)
	go func() {

		err = p.Con()
		if err != nil {
			err = utl.Msj.GetError("CN14")
			FinChan <- true
			return
		}
		sqltemp, args, err = p.NamedIn(query)
		if err != nil {
			err = utl.Msj.GetError("CN07")
			FinChan <- true
			return
		}

		filas, err = p.DBGO.Queryx(sqltemp, args...)
		if err != nil {
			p.Close()
			FinChan <- true
			return
		}
		if !indConect {
			p.Close()
		}

		FinChan <- true
		return
	}()

	for {
		select {
		default:
			continue
		case <-FinChan:

			if err != nil {

				return nil, err
			}

			return filas, nil
		}
	}

}

/*Query : Ejecuta un querie en la base de datos y
  devuelve un map dinamico para mostrar los datos donde le limitan la cantida
	de registro que debe de devolver
	indConect = true deja la conexion abierta
  Ejemplo:
  map[COD_CLI:50364481 NIS_RAD:5355046 SEC_NIS:1]
*/
func (p *StConect) Query(query StQuery, cantrow int, indConect bool) ([]StData, error) {
	var (
		err     error
		filas   *sqlx.Rows
		result  []StData
		args    []interface{}
		sqltemp string
	)
	if cantrow == 0 {
		return nil, utl.Msj.GetError("CN15")
	}

	FinChan := make(chan bool)
	defer close(FinChan)
	go func() {

		err = p.Con()
		if err != nil {
			FinChan <- true
			return
		}
		sqltemp, args, err = p.NamedIn(query)
		filas, err = p.DBGO.Queryx(sqltemp, args...)
		if err != nil {
			p.Close()
			FinChan <- true
			return
		}
		result, err = scanData(filas, cantrow)
		if err != nil {
			p.Close()
			filas.Close()
			err = utl.Msj.GetError("CN16")
			FinChan <- true
			return
		}

		if !indConect {
			p.Close()
		}
		filas.Close()
		FinChan <- true
		return
	}()

	for {
		select {
		default:
			continue
		case <-FinChan:

			if err != nil {

				return nil, err
			}

			return result, nil
		}
	}

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
	dato, err := p.Query(*prueba, 1, false)
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
	dato, err := p.Query(prueba, 1, false)
	if err != nil || len(dato) <= 0 {
		return false
	}
	num, erraux := dato[0].ToInt("REG")
	if num <= 0 || erraux != nil {
		return false
	}
	return true
}

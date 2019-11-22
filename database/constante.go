package database

const (
	/*Ora : conexion tipo oracle
	https://gopkg.in/rana/ora.v4
	*/
	Ora = "ORA"
	/*Post : conexion tipo postgres
	https://github.com/lib/pq
	*/
	Post = "POST"
	/*Mysql : conexion tipo mysql
	https://github.com/go-sql-driver/mysql
	*/
	Mysql = "MYSQL"
	/*Sqlser : conexion tipo sql server
	https://github.com/denisenkom/go-mssqldb
	*/
	Sqlser = "SQLSER"
	/*PrefixG : prefijo general para los parametros en las consultas ejemplo : :val1 oracle ? post  */
	PrefixG = ":n"
	/*INSERT : prefijo de insert */
	INSERT = "insert"
	/*UPDATE : prefijo de UPDATE */
	UPDATE = "update"
	/*DELETE : prefijo de DELETE */
	DELETE = "delete"
	/*SELECT : prefijo de select*/
	SELECT = "select"
	/*FROM : prefijo de tablas */
	FROM = "from"
)

var (
	/*CADCONN : contiene el formato de las cadenas de conexion*/
	CADCONN = map[string]string{
		Ora:    "%s/%s@%s:%d/%s",
		Post:   "postgres://%s:%s@%s:%d/%s?sslmode=%s",
		Mysql:  "%s:%s@tcp(%s:%d)/%s",
		Sqlser: "server=%s;user id=%s;password=%s;port=%d;database=%s;",
	}
	/*Prefijos : contiene  los diferentes prefijos de distintas base de datos.
	Ejemplo: select * from cliente where nombre=:n sera reemplazada por :val*/
	Prefijos = map[string]string{
		Ora:    ":val",
		Post:   "$",
		Mysql:  "?",
		Sqlser: ":",
	}
	/*PrefijosDB : contiene los string de conexion al momento de ejecutar la funcion open*/
	PrefijosDB = map[string]string{
		Ora:    "ora",
		Post:   "postgres",
		Mysql:  "mysql",
		Sqlser: "mssql",
	}
	/*Ssmodes : hace referencia si tienen conexion ssl
	0* disable - No SSL
	1* require - Always SSL (skip verification)
	2* verify-ca - Always SSL (verify that the certificate presented by the
	  server was signed by a trusted CA)
	3* verify-full - Always SSL (verify that the certification presented by
	  the server was signed by a trusted CA and the server host name
	  matches the one in the certificate)

	*/
	Ssmodes = []string{"disable", "require", "verify-ca", "verify-full"}
)

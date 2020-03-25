package test

import (
	"testing"

	db "github.com/rafael180496/libcore/database"
)

/*TestSqlLite : Se conecta a una base de datos  sqllite de configuracion*/
func TestSqlLite(t *testing.T) {
	var (
		conexion db.StConect
	)
	path := "config/sqllite.ini"
	t.Logf("Capturando path:%s", path)
	err := conexion.ConfigINI(path)
	if err != nil {
		t.Errorf("Error:%s", err.Error())
	}
	t.Logf("Conexion:%s", conexion.Conexion.ToString())
	t.Logf("Probando...")
	t.Logf("prueba:%v", conexion.Test())
}

/*TestPost : Se conecta a una base de datos  posgresql de configuracion*/
func TestPost(t *testing.T) {
	var (
		conexion db.StConect
	)
	path := "/home/rhidalgo/Documentos/go_projects/src/github.com/rafael180496/libcore/test/src/post.ini"
	t.Logf("Capturando path:%s", path)
	err := conexion.ConfigINI(path)
	if err != nil {
		t.Errorf("Error:%s", err.Error())
	}
	t.Logf("Conexion:%s", conexion.Conexion.ToString())
	t.Logf("Probando...")
	t.Logf("prueba:%v", conexion.Test())
	t.Logf("prueba:%v", conexion.Test())
	t.Logf("prueba:%v", conexion.Test())
}

/*TestPost : Se conecta a una base de datos  posgresql de configuracion con url*/
func TestPostURL(t *testing.T) {
	var (
		conexion db.StConect
	)
	URL := "POST:postgres/abc123/192.168.252.42:54320/([coreauth]-[]-[])"
	t.Logf("Capturando path:%s", URL)
	err := conexion.ConfigURL(URL)
	if err != nil {
		t.Errorf("Error:%s", err.Error())
	}
	t.Logf("Conexion:%s", conexion.Conexion.ToString())
	t.Logf("Probando...")
	t.Logf("prueba:%v", conexion.Test())
	t.Logf("prueba:%v", conexion.Test())
	t.Logf("prueba:%v", conexion.Test())
}

/*TestTestORAPost : Se conecta a una base de datos  ORACLE de configuracion*/
func TestORA(t *testing.T) {
	var (
		conexion db.StConect
	)
	path := "/home/rhidalgo/Documentos/go_projects/src/github.com/rafael180496/libcore/test/src/ora.json"
	t.Logf("Capturando path:%s", path)
	err := conexion.ConfigJSON(path)
	if err != nil {
		t.Errorf("Error:%s", err.Error())
	}
	t.Logf("Conexion:%s", conexion.Conexion.ToString())
	t.Logf("Probando...")
	t.Logf("prueba:%v", conexion.Test())
}

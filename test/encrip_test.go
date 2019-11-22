package test

import (
	"testing"

	"gitlab.com/gpctda/libcore/utility"
)

/*TestAes : prueba las funciones de encriptamiento aes */
func TestAes(t *testing.T) {
	key := "abc123"
	text := "hola mundo"
	t.Logf("Key:%s,Text:%s", key, text)
	bloque, err := utility.EncripAES(key, text)
	if err != nil {
		t.Errorf("Error:%s", err.Error())
	}
	t.Logf("BloqueEncrip:%s", bloque)
	bloque, err = utility.DesencripAES(key, bloque)
	t.Logf("BloqueDesencrip:%s", bloque)
}

/*TestMd5 : Genera un hash Md5 */
func TestMd5(t *testing.T) {
	key := "abc123"
	t.Logf("key:%s", key)
	t.Logf("bloquemd5:%s", utility.GeneredHashMd5(key))
}

/*Test256 : Genera un hash sha256 */
func Test256(t *testing.T) {
	key := "abc123"
	t.Logf("key:%s", key)
	t.Logf("bloque256:%s", utility.GeneredHashSha256(key))
}

/*TestToken : Genera un token con hash sha 256 unico */
func TestToken(t *testing.T) {
	key := "abc123"
	t.Logf("key:%s", key)
	t.Logf("bloquetoken:%s", utility.GenToken(key))
}

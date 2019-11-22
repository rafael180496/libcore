package test

import (
	"testing"

	"github.com/rafael180496/libcore/utility"
)

/*TestSendError : Envia un error con mensaje */
func TestSendError(t *testing.T) {
	t.Logf("Error:%s", utility.SendErrorCod("GE01").Error())
}

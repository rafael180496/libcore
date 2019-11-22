package test

import (
	"testing"

	"gitlab.com/gpctda/libcore/utility"
)

/*TestSendError : Envia un error con mensaje */
func TestSendError(t *testing.T) {
	t.Logf("Error:%s", utility.SendErrorCod("GE01").Error())
}

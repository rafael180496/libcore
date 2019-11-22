package utility

import (
	"fmt"
	"time"
)

type (
	/*StError : Estructura de manejo de error con mensajes personalizados*/
	StError struct {
		When time.Time
		What string
	}
)

/*Error : Muestra el mensaje personalizado del error*/
func (p *StError) Error() string {
	return fmt.Sprintf("%v: %v", time.Now(), p.What)
}

/*MandError : Envia un mensaje con la interfaz de error personalizado.*/
func (p *StError) MandError(mensaje string) error {
	p.What = mensaje
	p.When = time.Now()
	return p
}

/*SendErrorCod : envia un error con los mensajes guardados */
func SendErrorCod(cod string) error {
	return new(StError).MandError(Msj[cod])
}

package server

import "github.com/labstack/echo"

type (
	/*StInfoReq :  Obtiene  datos generales del request*/
	StInfoReq struct {
		HostOrig  string `json:"host"`
		IPRemote  string `json:"ip"`
		Browser   string `json:"browser"`
		SystemOI  string `json:"oi"`
		UserAgent string `json:"useragent"`
	}
	/*StDataEnv : Envio de datos generico*/
	StDataEnv struct {
		Code  int         `json:"code"`
		Error string      `json:"msj"`
		Data  interface{} `json:"data"`
	}
	/*HTTPTip : especifica el tipo de peticion a elaborar*/
	HTTPTip string
	/*Controller : estructura para crear Controladores */
	Controller struct {
		Pets []HTTPPet
	}

	/*HTTPPet Guarda una peticion HTTP basada en ECHO */
	HTTPPet struct {
		Path string
		Pet  func(echo.Context) error
		Tip  HTTPTip
	}
)

const (
	/*Peticiones http*/

	/*POST : peticion http POST*/
	POST HTTPTip = "POST"
	/*GET : peticion http GET*/
	GET HTTPTip = "GET"
	/*PUT : peticion http PUT*/
	PUT HTTPTip = "PUT"
	/*DELETE : peticion http DELETE*/
	DELETE HTTPTip = "DELETE"
)

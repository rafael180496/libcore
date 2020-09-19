package server

import echo "github.com/labstack/echo/v4"

type (
	/*ConfigServer : configuraciones del servicio generales */
	ConfigServer struct {
		/*Ipser : sirve para cuando este en modo produccion coloque la ip asignada al server*/
		Ipser string `ini:"ipser"`
		/*Debug : modo debug para el proyecto*/
		Debug bool `ini:"debug"`
		/*Local : habilita el modo local del servicio*/
		Local bool `ini:"local"`
		/*Puerto : Coloca el puerto del servicio*/
		Puerto int `ini:"puerto"`
		/*Protocol : protocolo del servicio*/
		Protocol string `ini:"protocol"`
	}
	/*StInfoPet : envia la informacion general de la peticion */
	StInfoPet struct {
		URL    string `json:"url"`
		Method string `json:"method"`
	}
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
	/*Tipos de protocolos : http https */

	/*HTTP : protocolo http*/
	HTTP = "http"
	/*HTTPS : protocolo https*/
	HTTPS = "https"
)

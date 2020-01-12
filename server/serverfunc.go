package server

import (
	utl "github.com/rafael180496/libcore/utility"

	"github.com/labstack/echo/v4"
	"github.com/mssola/user_agent"
)

/*ResultJSON : envia un json de los resultado del api rest */
func (p *StDataEnv) ResultJSON(e echo.Context) error {
	return e.JSON(p.Code, p)
}

/*ResultCodeJSON : envia un json de los resultado del api rest pero asignando el codigo */
func (p *StDataEnv) ResultCodeJSON(code int, e echo.Context) error {
	p.Code = code
	return e.JSON(p.Code, p)
}

/*ResultFullJSON : envia un json de los resultado del api rest pero asignando el codigo la data y el mensaje */
func (p *StDataEnv) ResultFullJSON(code int, msjerr string, data interface{}, e echo.Context) error {
	p.Code = code
	p.Error = msjerr
	p.Data = data
	return e.JSON(p.Code, p)
}

/*JSONSend : envia un json generic  */
func JSONSend(data StDataEnv, e echo.Context) error {
	return e.JSON(data.Code, data)
}

/*FindInfoReq : Envia la informacion de un request despues de una peticion */
func FindInfoReq(c echo.Context) (StInfoReq, error) {
	var (
		info StInfoReq
	)
	ua := user_agent.New(c.Request().UserAgent())
	info.UserAgent = c.Request().UserAgent()
	info.HostOrig = c.Scheme() + "://" + c.Request().Host
	nombrebrow, versionbrow := ua.Browser()
	info.IPRemote = c.RealIP()
	info.Browser = nombrebrow + " " + versionbrow
	info.SystemOI = ua.OS()
	return info, nil
}

/*Valid : Valida si la peticon es valida*/
func (p *HTTPPet) Valid() error {
	if !p.Tip.Valid() {
		return utl.Msj.GetError("SR01")
	}
	if utl.Trim(p.Path) == "" || utl.IsSpace(p.Path) {
		return utl.Msj.GetError("SR02")
	}
	if p.Pet == nil {
		return utl.Msj.GetError("SR03")
	}
	return nil
}

/*Valid : Valida si un tipo de perticion es valida.*/
func (p *HTTPTip) Valid() bool {
	switch *p {
	case POST, GET, PUT, DELETE:
		return true
	default:
		return false
	}
}

/*AsigServer : asigna peticiones a una instancia ECHO */
func AsigServer(e *echo.Echo, cs []Controller) error {
	for _, c := range cs {
		err := asigpet(e, c.Pets)
		if err != nil {
			return err
		}
	}
	return nil
}

/*asigpet : asigna las peticiones y las valida */
func asigpet(e *echo.Echo, ps []HTTPPet) error {
	for _, p := range ps {
		err := p.Valid()
		if err != nil {
			return err
		}
		err = findpet(e, p)
		if err != nil {
			return err
		}
	}
	return nil
}

/*findpet : busca la peticion para asignar*/
func findpet(e *echo.Echo, p HTTPPet) error {
	switch p.Tip {
	case POST:
		e.POST(p.Path, p.Pet)
		return nil
	case PUT:
		e.PUT(p.Path, p.Pet)
		return nil
	case GET:
		e.GET(p.Path, p.Pet)
		return nil
	case DELETE:
		e.DELETE(p.Path, p.Pet)
		return nil
	default:
		return utl.Msj.GetError("SR01")
	}
}

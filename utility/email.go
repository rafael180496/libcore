package utility

import (
	"bytes"
	"fmt"
	"mime/quotedprintable"
	"net/smtp"
	"strings"

	"github.com/asaskevich/govalidator"
)

type (
	/*StEmailAdmin : Estructura principal para envio de correro */
	StEmailAdmin struct {
		User  StAuthEmail `json:"user" valid:"-"`
		Dest  []string    `json:"dest" valid:"-"`
		Email StEmail     `json:"email" valid:"-"`
	}

	/*StAuthEmail  : credenciales para enviar correo */
	StAuthEmail struct {
		Email string `json:"email" valid:"email,required"`
		Pass  string `json:"pass" valid:"required,length(1|500)"`
		Host  string `json:"smtp" valid:"required,length(1|500)"`
		Port  string `json:"port" valid:"numeric,required"`
	}

	/*StEmail : estructura para enviar correo  */
	StEmail struct {
		HeadMsg     string `json:"head" valid:"required,length(1|500)"`
		BodyMsg     string `json:"body" valid:"required,length(1|10000)"`
		ContentType string `json:"content" valid:"required,length(1|500)"`
	}
)

/*Validar valida una estructura  StAuthEmail*/
func (p *StAuthEmail) Validar() bool {
	_, err := govalidator.ValidateStruct(p)
	return ReturnIf(err != nil, false, true).(bool)
}

/*Validar valida una estructura  StEmail*/
func (p *StEmail) Validar() bool {
	_, err := govalidator.ValidateStruct(p)
	return ReturnIf(err != nil, false, true).(bool)
}

/*Validar valida una estructura  StEmailAdmin*/
func (p *StEmailAdmin) Validar() bool {
	return ReturnIf(!p.User.Validar() || !p.Email.Validar() || len(p.Dest) <= 0, false, true).(bool)
}

/*SendMail : envia una cantidad X de correo masivo*/
func (p *StEmailAdmin) SendMail() error {
	host := p.User.Host
	hostfull := p.User.Host + ":" + p.User.Port
	err := smtp.SendMail(hostfull,
		smtp.PlainAuth("", p.User.Email, p.User.Pass, host),
		p.User.Email, p.Dest, []byte(p.ArmarEmail()))
	if err != nil {
		return err
	}
	return nil
}

/*AddUser : agrega un usuario para el envio de correo */
func (p *StEmailAdmin) AddUser(email, pass, host, port string) {
	p.User = NewAuth(email, pass, host, port)
}

/*AddDest : agrega un destinatario */
func (p *StEmailAdmin) AddDest(dest ...string) {
	clone := p.Dest
	clone = append(clone, dest...)
	p.Dest = clone
}

/*AddBody : agrega el cuerpo de un correo pudiendo agregar tambien html como texto plan */
func (p *StEmailAdmin) AddBody(content, header, body string) {
	p.Email.HeadMsg = header
	p.Email.ContentType = content
	p.Email.BodyMsg = body
}

/*ArmarEmail : arma el correo en general html o  texto plano*/
func (p *StEmailAdmin) ArmarEmail() string {
	header := make(map[string]string)
	header["From"] = p.User.Email
	receipient := ""
	receipient = strings.Join(p.Dest, ",")
	header["To"] = receipient
	header["Subject"] = p.Email.HeadMsg
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", p.Email.ContentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"
	message := ""
	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	var encodedMessage bytes.Buffer
	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(p.Email.BodyMsg))
	finalMessage.Close()
	message += "\r\n" + encodedMessage.String()
	return message
}

/*NewAuth : Crea una nueva instancia de un tipo Auth */
func NewAuth(email, pass, host, port string) StAuthEmail {
	return StAuthEmail{
		Email: email,
		Pass:  pass,
		Host:  host,
		Port:  port,
	}
}

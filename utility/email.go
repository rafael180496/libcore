package utility

type (
	/*StAuthEmail  : credenciales para enviar correo */
	StAuthEmail struct {
		Email string
		Pass  string
		Host  string
	}

	/*StEmail : estructura para enviar correo  */
	StEmail struct {
		User    []StAuthEmail
		Dest    []string
		HeadMsg string
		BodyMsg string
	}
)

/*AddUser : Agrega un usuario para ocuparlo de remitente*/
func (p *StEmail) AddUser(user StAuthEmail) error {
	return nil
}

/*NewAuth : Crea una nueva instancia de un tipo Auth */
func NewAuth(email, pass, host string) StAuthEmail {
	return StAuthEmail{
		Email: email,
		Pass:  pass,
		Host:  host,
	}
}

package correos

type CorreoGmail struct {}

func (receiver CorreoGmail) Enviar(mensg string)  bool {
	println("entro a enviar correo de GMAIL:",mensg)
	return true
}

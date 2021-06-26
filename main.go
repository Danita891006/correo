package main

import (
	msg "correo/src/Utils"
	correos "correo/src/correos"
	i "correo/src/strategy"
	"net/mail"
)

type Contexto struct {
	enviousMessage i.EnvioMessageInterface
}

func (c *Contexto) EnviarMail(mensg msg.Mensaje) bool {
	return c.enviousMessage.Enviar(mensg)
}

func main() {

	mensajeForGmail := msg.Mensaje{
		mail.Address{"Daniela Ortiz JAJAJ", "aleinadoh89@gmail.com"},
		mail.Address{"JC Fuentes ", "carkar@gmail.com"},
		"Envio para gmail"}

	mensajeForVantilabs := msg.Mensaje{
		mail.Address{"Daniela Ortiz", "aleinadoh89@gmail.com"},
		mail.Address{"vanti jc ", "juan.fuentes@vantilabs.com"},
		"envío para vantilabs"}

	x := Contexto{correos.CorreoGmail{}}
	x.EnviarMail(mensajeForGmail)

	j := Contexto{correos.CorreoGmail{}}
	j.EnviarMail(mensajeForVantilabs)
}

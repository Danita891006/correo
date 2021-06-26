package factory
import (
	m "../strategy"
)


type Contexto struct {
	correo m.EnvioMensaje
}

func (c *Contexto) EnviarMail(mensg string) bool {
	return c.correo.Enviar(mensg)
}


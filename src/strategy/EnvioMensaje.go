package strategy

import msg "correo/src/Utils"

type EnvioMessageInterface interface {
	Enviar(msg.Mensaje) bool
}

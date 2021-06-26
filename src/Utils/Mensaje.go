package Utils

import "net/mail"

type Mensaje struct {
	From    mail.Address
	To      mail.Address
	Subject string
}

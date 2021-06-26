package correos

import (
	"bytes"
	msg "correo/src/Utils"
	"crypto/tls"
	"fmt"
	"html/template"
	"net/smtp"
)

type CorreoGmail struct{}
type Client struct {
	Nombre string
}

func (receiver CorreoGmail) Enviar(mensg msg.Mensaje) bool {
	fmt.Printf("mensage desde gmail :%s", mensg)

	headers := make(map[string]string)
	headers["From"] = mensg.From.String() //from.String()
	headers["To"] = mensg.To.String()     //to.String()
	headers["Subject"] = mensg.Subject
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	clientName := Client{Nombre: mensg.To.Name}

	message := ""

	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	t, err := template.ParseFiles("html/template-correo.html")
	if err != nil {
		fmt.Printf("Error al obtener html%d\n", err.Error())
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, clientName)
	message += buf.String()

	servername := "smtp.gmail.com:465"
	host := "smtp.gmail.com"

	auth := smtp.PlainAuth("", "aleinadoh89@gmail.com", "hzulmsdbkryblqbw", host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsConfig)
	if err != nil {
		fmt.Printf("Error al configurar tcp%d\n", err.Error())
		return false
	}

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		fmt.Printf("Error al crear el cliente%d\n", err.Error())
		return false
	}

	err = client.Auth(auth)
	if err != nil {
		fmt.Printf("Error al crear autorizacion%d\n", err.Error())
		return false
	}

	err = client.Mail(mensg.From.Address)
	if err != nil {
		fmt.Printf("Error al obtener from%d\n", err.Error())
		return false
	}

	err = client.Rcpt(mensg.To.Address)
	if err != nil {
		fmt.Printf("Erro al obtener to%d\n", err.Error())
		return false
	}

	w, err := client.Data()
	if err != nil {
		fmt.Printf("Error al obtener el mensaje%d\n", err.Error())
		return false
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		fmt.Printf("Error al escribir el mensaje%d\n", err.Error())
		return false
	}

	err = w.Close()
	if err != nil {
		fmt.Printf("Error al cerrar exritura%d\n", err.Error())
		return false
	}

	client.Quit()

	return true
}

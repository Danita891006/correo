package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net/mail"
	"net/smtp"
)

type Client struct {
	Nombre string
}
func main()  {
	from := mail.Address{"Daniela Ortiz","aleinadoh89@gmail.com"}
	to := mail.Address{"Juan Carlos Fuentes","juan.fuentes@vantilabs.com"}
	subject := "Prueba envi\u00F3 correo"

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = `text/html; charset="UTF-8"`
	clientName := Client{ Nombre : to.Name}

	message := ""

	for k, v := range headers{
		message += fmt.Sprintf("%s: %s\r\n",k,v)
	}

	t ,err := template.ParseFiles("html/template-correo.html")
	if err != nil{
		fmt.Printf("Error al obtener html%d\n",err.Error())
	}

	buf := new (bytes.Buffer)

	err = t.Execute(buf, clientName)
	message += buf.String()

	servername := "smtp.gmail.com:465"
	host := "smtp.gmail.com"

	auth :=  smtp.PlainAuth("","aleinadoh89@gmail.com","hzulmsdbkryblqbw",host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName: host,
	}

	conn,err := tls.Dial("tcp", servername, tlsConfig)
	if err != nil{
		fmt.Printf("Error al configurar tcp%d\n",err.Error())
	}


	client,err := smtp.NewClient(conn,host)
	if err != nil{
		fmt.Printf("Error al crear el cliente%d\n",err.Error())
	}


	err = client.Auth(auth)
	if err != nil{
		fmt.Printf("Error al crear autorizacion%d\n",err.Error())
	}


	err = client.Mail(from.Address)
	if err != nil{
		fmt.Printf("Error al obtener from%d\n",err.Error())
	}


	err = client.Rcpt(to.Address)
	if err != nil{
		fmt.Printf("Erro al obtener to%d\n",err.Error())
	}


	w, err :=client.Data()
	if err != nil{
		fmt.Printf("Error al obtener el mensaje%d\n",err.Error())
	}


	_,err = w.Write([]byte(message))
	if err != nil{
		fmt.Printf("Error al escribir el mensaje%d\n",err.Error())
	}


	err = w.Close()
	if err != nil{
		fmt.Printf("Error al cerrar exritura%d\n",err.Error())
	}


	client.Quit()
}

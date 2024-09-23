package demo

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type Mail struct {
	Sender  string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}

type SMTPServer struct {
	Host      string
	Port      string
	TLSConfig *tls.Config
}

func EmailController(mail Mail) {
	fmt.Println("mail: ", mail)
	send(mail)
}

func (s *SMTPServer) ServerName() string {
	return s.Host + ":" + s.Port
}

func send(mail Mail) {
	messageBody := mail.BuildMessage()
	fmt.Printf("messageBody: %v\n", messageBody)

	smtpServer := SMTPServer{
		Host: "smtp.gmail.com",
		Port: "465",
	}
	smtpServer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.Host,
	}

	_ = "arjesh.vadadoriya@grewon.com"
	password := "oltx mili uugf ffiz"
	auth := smtp.PlainAuth("", mail.Sender, password, smtpServer.Host)

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), smtpServer.TLSConfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.Host)
	if err != nil {
		log.Panic(err)
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// step 2: add all from and to
	if err = client.Mail(mail.Sender); err != nil {
		log.Panic(err)
	}
	receivers := append(mail.To, mail.Cc...)
	receivers = append(receivers, mail.Bcc...)
	for _, k := range receivers {
		log.Println("sending to: ", k)
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	// Data
	wr, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = wr.Write([]byte(messageBody))
	if err != nil {
		log.Panic(err)
	}

	err = wr.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Mail sent successfully")
}

func (m *Mail) BuildMessage() string {
	header := ""
	header += fmt.Sprintf("From: %s\r\n", m.Sender)
	if len(m.To) > 0 {
		header += fmt.Sprintf("To: %s\r\n", strings.Join(m.To, ";"))
	}
	if len(m.Cc) > 0 {
		header += fmt.Sprintf("Cc: %s\r\n", strings.Join(m.Cc, ";"))
	}
	header += fmt.Sprintf("Subject: %s\r\n", m.Subject)
	header += "\r\n" + m.Body

	return header
}

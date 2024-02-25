package handlers

import (
	"crypto/tls"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/AthirsonSilva/music-streaming-api/cmd/server/models/dto"
	mail "github.com/xhit/go-simple-mail/v2"
)

var EmailChannel = make(chan dto.EmailData)
var (
	host     = os.Getenv("SMTP_HOST")
	port     = os.Getenv("SMTP_PORT")
	password = os.Getenv("SENDER_EMAIL_PASSWORD")
	from     = os.Getenv("SENDER_EMAIL_ADDRESS")
)

func ListenForEmail() {
	func() {
		log.Println("Listening for email...")
		for {
			emailData := <-EmailChannel
			log.Printf("Sending email to: %s\n", emailData.To)
			SendSimpleEmailMessage(emailData)
		}
	}()
}

func SendSimpleEmailMessage(emailData dto.EmailData) {
	log.Println("Sending email...")

	port, err := strconv.Atoi(port)
	if err != nil {
		log.Printf("Failed to convert port: %s\n", err)
		return
	}

	log.Printf("Setting up smtp server: %s:%d\n", host, port)

	server := mail.NewSMTPClient()
	server.Host = host
	server.Port = port
	server.Username = from
	server.Password = password
	server.KeepAlive = false
	server.ConnectTimeout = 20 * time.Second
	server.SendTimeout = 20 * time.Second
	server.Encryption = mail.EncryptionTLS

	log.Println("Connecting to smtp server...")

	client, err := server.Connect()
	if err != nil {
		log.Printf("Failed to connect to smtp server: %s\n", err)
		return
	}

	log.Println("Connected to smtp server! Sending email...")

	email := mail.NewMSG()
	email.SetFrom(from)
	email.AddTo(emailData.To)
	email.SetSubject(emailData.Subject)
	email.SetBody(mail.TextHTML, emailData.Body)

	err = email.Send(client)
	if err != nil {
		log.Printf("Failed to send mail: %s\n", err)
		return
	}

	log.Println("Email sent!")
}

func SendVerificationEmail(emailData dto.EmailData) {
	log.Println("Sending email...")

	authenticate := smtp.PlainAuth("", from, password, host)
	tlsConfigurations := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", host+":"+port, tlsConfigurations)
	if err != nil {
		log.Printf("Failed to connect to smtp server: %s\n", err)
		return
	}

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Printf("Failed to create smtp client: %s\n", err)
		return
	}

	if err = client.Auth(authenticate); err != nil {
		log.Printf("Failed to authenticate: %s\n", err)
		return
	}

	if err = client.Mail(from); err != nil {
		log.Printf("Failed to send mail: %s\n", err)
		return
	}

	if err = client.Rcpt(emailData.To); err != nil {
		log.Printf("Failed to send mail: %s\n", err)
		return
	}

	writer, err := client.Data()
	if err != nil {
		log.Printf("Failed to send mail: %s\n", err)
		return
	}

	emailMessage := "Subject: " + emailData.Subject + "\r\n\r\n" + emailData.Body
	_, err = writer.Write([]byte(emailMessage))
	if err != nil {
		log.Printf("Failed to send mail: %s\n", err)
		return
	}

	if err = writer.Close(); err != nil {
		log.Printf("Failed to send mail: %s\n", err)
		return
	}

	if err = client.Quit(); err != nil {
		log.Printf("Failed to send mail: %s\n", err)
		return
	}

	log.Println("Email sent successfully")
}

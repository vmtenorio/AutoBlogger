// Package mailutils provides several functions to build and send the email
// to the email list provided
package mailutils

import (
  "fmt"
  "net/smtp"
  fu "autoBlogger/fileutils"
  "net/mail"
  "crypto/tls"
  "log"
  "errors"
)


func SendEmail (content string, senderEmail string, senderPassword string, sett fu.Settings) bool {
//  to := []string{sett.MailList}
  from := mail.Address{"Even Deeper Learning", senderEmail}
  to   := mail.Address{"Even Deeper Learning Group", sett.MailList}

  smtpHost := sett.SmtpServer
  smtpPort := sett.SmtpPort

  // Setup headers
  headers := make(map[string]string)
  headers["From"] = from.String()
  headers["To"] = to.String()
  headers["Subject"] = "Your 4 Minutes ML Pills"
  headers["Content-Type"] = "text/html; charset=\"ISO-8859-1\";"

  // Setup message
  message := ""
  for k,v := range headers {
    message += fmt.Sprintf("%s: %s\r\n", k, v)
  }
  message += "\r\n" + content

  // Connect to the SMTP Server
  auth := LoginAuth(senderEmail, senderPassword)
  // The empty string is for the identity argument. As it is empty, it uses the username

  // TLS config
  tlsconfig := &tls.Config {
    InsecureSkipVerify: true,
    ServerName: smtpHost,
  }

  c, err := smtp.Dial(smtpHost + ":" + smtpPort)
  if err != nil {
    log.Panic(err)
  }

  if err = c.StartTLS(tlsconfig); err != nil {
    log.Panic(err)
  }

  // Auth
  if err = c.Auth(auth); err != nil {
    log.Panic(err)
  }

  // To && From
  if err = c.Mail(senderEmail); err != nil {
    log.Panic(err)
  }

  if err = c.Rcpt(sett.MailList); err != nil {
    log.Panic(err)
  }

  // Data
  w, err := c.Data()
  if err != nil {
    log.Panic(err)
  }

  _, err = w.Write([]byte(message))
  if err != nil {
    log.Panic(err)
  }

  err = w.Close()
  if err != nil {
    log.Panic(err)
  }

  c.Quit()
  return true

}

// Login Auth
type loginAuth struct {
  username, password string
}

func LoginAuth(username, password string) smtp.Auth {
  return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
  return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
  if more {
    switch string(fromServer) {
    case "Username:":
      return []byte(a.username), nil
    case "Password:":
      return []byte(a.password), nil
    default:
      return nil, errors.New("Unkown fromServer")
    }
  }
  return nil, nil
}

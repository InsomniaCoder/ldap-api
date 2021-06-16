package mail

import (
	"errors"
	"fmt"
	"net/smtp"
)

func SendEmail(userId, userPassword, userEmail string) (err error) {

	auth := LoginAuth("", "")
	to := []string{userEmail}
	msg := []byte(fmt.Sprintf("Your LDAP Account is %s and your password is %s", userId, userPassword))
	err = smtp.SendMail("smtp-mail.outlook.com:587", auth, "sender-email", to, msg)
	return
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
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

package ldap

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/go-ldap/ldap"
	"github.com/sethvargo/go-password/password"
)

func CreateLDAPAccount(userId, firstName, lastName, userEmail string) (generatedPassword string, ldapErr error) {

	ldapURL := os.Getenv("LDAP_URL")
	// connect to LDAP

	l, err := ldap.DialURL(ldapURL)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer l.Close()

	log.Println("connected")

	// authenticate to LDAP
	adminDN := os.Getenv("ADMIN_DN")
	adminPasswd := os.Getenv("ADMIN_PASSWD")

	err = l.Bind(adminDN, adminPasswd)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	log.Println("authenticated")

	// TODO check user exists

	// add New User

	// Generate a password that is 16 characters long with 3 digits, 3 symbols,
	// allowing upper and lower case letters, disallowing repeat characters.
	userPassword, err := password.Generate(16, 3, 3, false, false)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	// -h {SSHA} is the default
	hashedUserPassword, _ := exec.Command("slappasswd", "-s", userPassword).Output()

	log.Printf("generated password is %s", userPassword)
	log.Printf("hashed generated password is %x", hashedUserPassword)

	userDNFormat := os.Getenv("USER_DN_FORMAT")
	userDN := fmt.Sprintf(userDNFormat, userId)

	addReq := ldap.NewAddRequest(userDN, []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "organizationalPerson", "inetOrgPerson", "person", "posixAccount"})
	addReq.Attribute("cn", []string{userId})
	addReq.Attribute("sn", []string{lastName})
	addReq.Attribute("givenName", []string{firstName})
	addReq.Attribute("uid", []string{userId})
	addReq.Attribute("userPassword", []string{string(hashedUserPassword)})
	addReq.Attribute("mail", []string{userEmail})
	addReq.Attribute("uidNumber", []string{"10009"})
	addReq.Attribute("gidNumber", []string{"10000"})
	addReq.Attribute("homeDirectory", []string{fmt.Sprintf("/home/%s", userId)})
	addReq.Attribute("loginShell", []string{"/bin/bash"})

	if err := l.Add(addReq); err != nil {
		log.Fatal("error adding service:", addReq, err)
		return "", err
	}

	log.Println(fmt.Sprintf("user %s has been created", userId))
	return userPassword, nil
}

func ResetPassword(userId, newPassword string) (changedPassword string, ldapErr error) {

	ldapURL := os.Getenv("LDAP_URL")
	// connect to LDAP

	l, err := ldap.DialURL(ldapURL)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer l.Close()

	log.Println("connected")

	// authenticate to LDAP
	adminDN := os.Getenv("ADMIN_DN")
	adminPasswd := os.Getenv("ADMIN_PASSWD")

	err = l.Bind(adminDN, adminPasswd)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	log.Println("authenticated")

	// TODO check user exists

	// Reset Password
	// -h {SSHA} is the default
	hashedUserPassword, _ := exec.Command("slappasswd", "-s", newPassword).Output()

	log.Printf("changing password is %s", newPassword)
	log.Printf("hashed changing password is %x", hashedUserPassword)

	userDNFormat := os.Getenv("USER_DN_FORMAT")
	userDN := fmt.Sprintf(userDNFormat, userId)

	modifyRequest := ldap.NewModifyRequest(userDN, []ldap.Control{})
	modifyRequest.Replace("userPassword", []string{string(hashedUserPassword)})

	if err := l.Modify(modifyRequest); err != nil {
		log.Fatal("error adding service:", modifyRequest, err)
		return "", err
	}

	log.Println(fmt.Sprintf("user %s password has been modified", userId))
	return newPassword, nil
}

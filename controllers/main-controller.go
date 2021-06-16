package controllers

import (
	"fmt"
	"net/http"

	"github.com/insomniacoder/ldap-api/ldap"
	"github.com/insomniacoder/ldap-api/models"

	"github.com/gin-gonic/gin"
)

//Create LDAP account
func CreateNewLDAPUser(c *gin.Context) {

	var createRequest models.CreateRequest
	c.BindJSON(&createRequest)

	fmt.Printf("creating LDAP account with detail %s", createRequest)

	generatedPassword, ldapErr := ldap.CreateLDAPAccount(createRequest.ID, createRequest.FirstName, createRequest.LastName, createRequest.Email)

	if ldapErr != nil {
		c.JSON(http.StatusInternalServerError, ldapErr)
	}

	defer c.JSON(http.StatusOK, generatedPassword)
}

// func sendEmail() {
// 	// send email when we have correct username / password
// 	if err := mail.SendEmail("id", "pwd", "tanat.l@monix.co.th"); err != nil {
// 		log.Fatal("send email fail ", err)
// 	}
// }

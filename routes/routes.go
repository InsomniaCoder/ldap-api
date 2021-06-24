package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/insomniacoder/ldap-api/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	ldap := r.Group("/ldaps")
	{
		ldap.POST("users", controllers.CreateNewLDAPUser)
		ldap.PUT("users/password", controllers.ResetUserLDAPPassword)
	}
	return r
}

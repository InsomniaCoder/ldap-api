package main

import (
	"github.com/insomniacoder/ldap-api/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run()
}

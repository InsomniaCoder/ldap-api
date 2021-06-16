# Instruction

This project is created to provide RESTFul APIs for LDAP operations.

It is created using `go-ldap` package and wrapped in `gin-gonic` for providing API.

## Local

```
go mod init github.com/insomniacoder/ldap-api

//set up environment varialbe in .envrc
// LDAP_URL, ADMIN_DN, ADMIN_PASSWD, USER_DN_FORMAT
// and run direnv allow .
go run main.go
```


## API usage

POST /ldaps/users

```
{
 "id"   : "testgo",
 "firstName"   : "firstName",
 "lastName"   : "lastName",
 "email"   : "testgo3@mail.com"
}
```
you should get response as `password`


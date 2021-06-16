# Instruction

This project is created to provide RESTFul APIs for LDAP operations.

It is created using `go-ldap` package and wrapped in `gin-gonic` for providing API.

## Local


- `go mod init github.com/insomniacoder/ldap-api`
- set up environment varialbe in .envrc
```
export LDAP_URL="ldap://<your-ldap-domain>:<port>"
export ADMIN_DN="cn=admin,dc=your,dc=domain,dc=com"
export ADMIN_PASSWD="password"
export USER_DN_FORMAT="uid=%s,ou=personal,dc=your,dc=domain,dc=com"
```
- and run `direnv allow .`
- `go run main.go`


## API usage

### Endpoints

#### POST /ldaps/users

request body:

```
{
 "id"   : "testgo",
 "firstName"   : "firstName",
 "lastName"   : "lastName",
 "email"   : "testgo3@mail.com"
}
```
response:

`generated password`


## How To
Add company:
```sh
curl -X POST -H "Content-Type: application/json" --data "{\"code\": \"J1234\", \"name\": \"avata\", \"country\": \"Fidji\", \"website\": \"avata.fj\", \"phone\": \"+55 12345\"}" http://localhost:3000/api/v1/company
```
Get company with ID = 1:
```sh
curl http://localhost:3000/api/v1/company/1
```
Get all companies:
```sh
curl http://localhost:3000/api/v1/company/
```
Update company with ID = 1, change in field code:
```sh
curl -X PUT -H "Content-Type: application/json" --data "{\"id\": 1,\"code\": \"Jxxxx\", \"name\": \"avata\", \"country\": \"Fidji\", \"website\": \"avata.fj\", \"phone\": \"+55 12345\"}" http://localhost:3000/api/v1/company
```
See also: 
```html
https://stackoverflow.com/questions/39333102/how-to-create-or-update-a-record-with-gorm
```
Delete company with ID = 1:
```sh
curl -X DELETE http://localhost:3000/api/v1/company/1
```


## Leftovers
### Move database operations to context
### Improve error handling
example: send same creation twice
### Setup golangci linter
### Check or improve the memory allignment of structs
### Move database connection as singleton
### Split into packages
### Configuration load
### Create interface for repo
### Improve app shutdown
### Diminish exported objects / variables
### Model validation
### Logging
### Assess for SQL injection
### Move to constants error messages
### Less Code duplication
### IP service
see:
```html
https://github.com/gofiber/recipes/tree/master/geoip
```
### Context timeout for IP service
### Fiber protection middlewares


## Resources
```html
https://tutorialedge.net/golang/basic-rest-api-go-fiber/
https://www.moesif.com/blog/technical/api-design/Which-HTTP-Status-Code-To-Use-For-Every-CRUD-App/
```
# Golang JWT Server

This is a simple jwt authentication server written in Golang with the Echo framework. 
Much of the code came from Echo framework's tutorial on JWT : https://echo.labstack.com/middleware/jwt

Echo Framework was chosen for its lightweight and simplicity. 
Go Mod is used for dependency management (https://blog.golang.org/using-go-modules)

## PREREQS
- Go
- Docker 

### How to run?
This will trigger the Makefile to run go test, build a cointainer and run the server exposed on port :1323 
```
$ make
```
### Manual Testing
```
$ curl -X GET localhost:1323/
$ curl -X POST -d 'username=jon' -d 'password=shhh!' localhost:1323/login
```
Once you have logged in with those credentials specifically, you can access the private route: 
```
$ curl localhost:1323/restricted -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9uIFNub3ciLCJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTY4MTc4MTQzfQ.a_S9ti_CT4VWyrLpOAAo6dVXshI0N9pZjiGfI3s_H6E"
```


### Improvements
- Write more functions within the Login handler so it is smaller and easily testable
- Write more tests especially for /restricted endpoint as well as the expiry times for each token. 
- CI/CD deployment 
#Golang JWT Server

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

### Improvements
- Write more functions within the Login handler so it is smaller and easily testable
- Write more tests ( unit and integration ) 
- CI/CD deployment 
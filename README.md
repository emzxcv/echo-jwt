# Golang JWT Server

This is a simple jwt authentication server written in Golang with the Echo framework. It is able to set custom claims.
Much of the code came from Echo framework's tutorial on JWT : https://echo.labstack.com/middleware/jwt

Echo Framework was chosen for its lightweight and simplicity. 
Go Mod is used for dependency management (https://blog.golang.org/using-go-modules)

## PREREQS
- Go
- Docker 

### How to run?
This will trigger the Makefile to run go test, build a container and run the server exposed on port :1323
```
$ make
```
### Manual Testing
Test the server is up 
```
$ curl -X GET localhost:1323/
```

Test you can get an access token (1 day expiry) & a refresh token (7 days expiry)
```
$ curl -X POST -d 'username=jon' -d 'password=shhh!' localhost:1323/login
```
You should expect to see this reponse:
```
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9uIFNub3ciLCJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTY4Njk4MjY5LCJzdWIiOiIxIn0.LekFS1tc0hoTJMv_EPkoUEP7tKFT-VUzHANiH09f7Jw",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9uIFNub3ciLCJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTY4MTc5ODY5fQ.rnOv0MGrDSUW_h0Q2W9hzUdpKG-zCQdxAMPrLiAtld0"
}
```

Once you have logged in with those credentials specifically, you can access the private route replacing the {{TOKEN}} with either the token or refresh_token from previous request: 
```
$ curl localhost:1323/restricted -H "Authorization: Bearer {{TOKEN}}"
```

### Improvements
- Write more functions within the Login handler so it is smaller and easily testable. Eg. Write function to set custom claims. 
- Write more tests especially for /restricted endpoint as well as the expiry times for each token. 
- CI/CD deployment 
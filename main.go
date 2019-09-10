package main

import (
	"github.com/dgrijalva/jwt-go"
	auth "github.com/emzxcv/echo-jwt/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

// User
type User struct {
	Username string `json:"username" `
	Password string `json:"password" `
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// create handler
	h := &auth.Handler{}

	// Login route
	e.POST("/login", h.login)

	// Unauthenticated route
	e.GET("/", h.accessible)

	// Restricted group
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", h.restricted)

	e.Logger.Fatal(e.Start(":1323"))
}

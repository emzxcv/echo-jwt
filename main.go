package main

import (
	auth "github.com/emzxcv/echo-jwt/handler"
	m "github.com/emzxcv/echo-jwt/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// create handler
	h := &auth.Handler{}

	// Login route
	e.POST("/login", h.Login)

	// Unauthenticated route
	e.GET("/", h.Accessible)

	// Restricted group
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &m.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", h.Restricted)

	e.Logger.Fatal(e.Start(":1323"))
}

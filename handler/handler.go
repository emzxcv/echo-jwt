package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	m "github.com/emzxcv/echo-jwt/models"
	"github.com/labstack/echo"
)

type Handler struct{}

func (h *Handler) Login(c echo.Context) error {
	u := new(m.User)
	if err := c.Bind(u); err != nil {
		return echo.ErrBadRequest
		// Check in your db if the user exists or not
	}

	// Throws unauthorized error
	if u.Username != "jon" || u.Password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &m.JwtCustomClaims{
		"Jon Snow",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	// Set custom claims
	rtClaims := &m.JwtCustomClaims{
		"Jon Snow",
		true,
		jwt.StandardClaims{
			Subject:   "1",
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}

	// Generate encoded refrest token and send it as response.
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token":         t,
		"refresh_token": rt,
	})
}

func (h *Handler) Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible. Please navigate to /login or /restricted.")
}

func (h *Handler) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*m.JwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

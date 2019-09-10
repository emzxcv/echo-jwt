package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type Handler struct{}

func (h *Handler) login(c echo.Context) error {
	// username := c.FormValue("username")
	// password := c.FormValue("password")
	u := new(User)
	if err := c.Bind(u); err != nil {
		return echo.ErrBadRequest
		// Check in your db if the user exists or not
	}

	// Throws unauthorized error
	if u.Username != "jon" || u.Password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
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
	rtClaims := &jwtCustomClaims{
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

func (h *Handler) accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible. Please navigate to /login or /restricted.")
}

func (h *Handler) restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

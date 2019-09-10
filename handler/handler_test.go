package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON        = `{"username":"jon","password":"shhh!"}`
	badUserNameJSON = `{"username":"jonnn","password":"sh!"}`
	badPasswordJSON = `{"username":"jon","password":"shhhhhhhhh!"}`
	badFormatJSON   = `{"name":"jon","password":"shhh!"}`
)

type JwtResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func TestLoginSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{}

	// Assertions
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "token")
		assert.Contains(t, rec.Body.String(), "refresh_token")
	}
}
func TestBadUsername(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(badUserNameJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{}

	// Assertions
	assert.EqualError(t, h.Login(c), "code=401, message=Unauthorized")
}

func TestBadPassword(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(badPasswordJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{}

	// Assertions
	assert.EqualError(t, h.Login(c), "code=401, message=Unauthorized")
}

func TestBadFormatJSON(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(badFormatJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{}

	// Assertions
	assert.EqualError(t, h.Login(c), "code=401, message=Unauthorized")
}

func TestEndpointAccessible(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{}

	// Assertions
	if assert.NoError(t, h.Accessible(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Accessible")
	}
}

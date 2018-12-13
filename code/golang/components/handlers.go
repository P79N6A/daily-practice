package components

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/color"
)

const (
	secretKey = "Hw&%@#hwt+k_zt_4%#@"
)

type (
	// ColoredFunc func(msg interface{}, styles ...string) string
	ColoredFunc func(msg interface{}, styles ...string) string

	jwtCustomClaims struct {
		Name string `json:"name"`
		jwt.StandardClaims
	}

	// ValidateFunc func(username, password string) bool
	ValidateFunc func(username, password string) bool

	// Message within JSON
	Message struct {
		Message string `json:"message"`
	}
)

// JSONWithError return a JSON error message
func JSONWithError(ctx *echo.Context, statusCode int, err *error) error {
	return (*ctx).JSON(statusCode, &Message{(*err).Error()})
}

// PrintlnWithClr (c ColoredFunc, msg interface{})
func PrintlnWithClr(c ColoredFunc, msg interface{}) {
	color.Println(c(msg))
}

func GenHandler(msg string) func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, msg)
	}
}

func PrintlnWithIndent(spaces int, args ...interface{}) {
	vargs := append(append([]interface{}{}, strings.Repeat(" ", spaces)), args...)
	log.Println(vargs...)
}

// RainbowCrap (msg string)
func RainbowCrap(msg string) {
	PrintlnWithClr(color.Green, msg)
	PrintlnWithClr(color.Red, msg)
	PrintlnWithClr(color.Yellow, msg)
	PrintlnWithClr(color.Blue, msg)
	PrintlnWithClr(color.Magenta, msg)
	PrintlnWithClr(color.Cyan, msg)
	PrintlnWithClr(color.Grey, msg)
}

// StreamingHandler 长连接
func StreamingHandler(c echo.Context) error {
	res := c.Response()
	gone := res.CloseNotify()
	res.Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	res.WriteHeader(http.StatusOK)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	fmt.Fprint(res, "<pre><strong>ℷ.u.me ⟹ ridiculous (LOL</strong>\n\n<code>")
	for {
		fmt.Fprintf(res, "%v\n", time.Now())
		res.Flush()
		select {
		case <-ticker.C:
		case <-gone:
			break
		}
	}
}

// JSONPHandler ...
func JSONPHandler(c echo.Context) error {
	callback := c.QueryParam("callback")
	content := struct {
		Response string    `json:"response"`
		Version  time.Time `json:"version"`
	}{
		Response: "Holy crap!",
		Version:  time.Now().UTC(),
	}
	return c.JSONP(http.StatusOK, callback, &content)
}

// JWTLogin ...
func JWTLogin(expires time.Duration, f ValidateFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		if f(username, password) {
			claims := &jwtCustomClaims{
				username,
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(expires).Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			t, err := token.SignedString([]byte(secretKey))
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, echo.Map{
				"token": t,
			})
		}

		return echo.ErrUnauthorized
	}
}

// JWTMiddleware ...
func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte(secretKey),
	}
	return middleware.JWTWithConfig(config)
}

// BasicAuthHandler ...
func BasicAuthHandler(username, password string) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(u, p string, c echo.Context) (success bool, err error) {
		if u == username && p == password {
			success = true
		}
		return success, nil
	})
}

package main

import (
	"runtime/pprof"
	"time"

	comp "./components"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	port = ":1323"
)

func main() {
	comp.StartProfile("./main.prof")
	defer pprof.StopCPUProfile()

	// SOME CRAP...
	comp.RainbowCrap("Server running on" + port)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} [${latency_human}] ${status} ${method} ${uri}` + "\n",
	}))
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		comp.PrintlnWithIndent(3, "middleware")
		return h
	})
	e.Pre(func(h echo.HandlerFunc) echo.HandlerFunc {
		comp.PrintlnWithIndent(1, "pre-middleware")
		return h
	})

	e.GET("/", comp.GenHandler("Luke, are u ready?"))
	e.GET("/users/:uid", comp.GenHandler("Luke, let me in."))
	e.GET("/papa/*", comp.GenHandler("Luke, I'm ur papa."))
	e.GET("/v2/v3/halo", comp.GenHandler("Luke, halo."))
	e.GET("/v2/v3/hello", comp.GenHandler("Luke, hello."))
	e.GET("/v2/v3/holy", comp.GenHandler("Luke, holy crap."))

	// long connection
	e.GET("/stream", comp.StreamingHandler)
	// JSONP
	e.GET("/jsonp", comp.JSONPHandler)

	// BASIC AUTH
	e.GET("/basic", comp.GenHandler("hi, basic auth"), comp.BasicAuthHandler("holy", "crap"))

	// JWT
	e.POST("/login", comp.JWTLogin(20*time.Second, func(username, password string) bool {
		return username == "holy" && password == "crap"
	}))
	g := e.Group("/admin")
	g.Use(comp.JWTMiddleware())
	g.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		comp.PrintlnWithIndent(5, "admin middleware")
		return func(c echo.Context) error {
			comp.PrintlnWithIndent(7, c.Scheme(), c.Path())
			return h(c)
		}
	})
	g.GET("/crap", comp.GenHandler("admin, crap"))
	g.GET("/quack", comp.GenHandler("admin, quack"))

	// STATIC
	e.Static("/static", "assets")

	// e.TraverseTree()

	// e.Logger.Fatal(e.StartTLS(":1323", "certs/cert.pem", "certs/key.pem"))
	e.Logger.Fatal(e.Start(port))
}

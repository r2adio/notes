package server

import (
	"net/http"

	"goth/cmd/web"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	//logs every http request; Ist: captures all requests, even those that panic
	e.Use(middleware.Logger())
	//catches panics and returns 500 error instead of crashing; and protects server crashes from unhandled panics
	//protects all subsequent middleware and handlers
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//allows any origin; security risk: very permissive for production; specify exact origins, ["https://urdomain.com"]
		AllowOrigins: []string{"https://*", "http://*"},
		//covers all RESTful operations; OPTIONS: req. for preflight requests; PATCH: for partial updates
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		//Authorization: for JWT tokens and API keys; Content-Type: required for json req.; X-CSRF-Token: CSRF protecttion
		AllowHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		//allows cookies and auth headers; cant use with wildcard origins in production; session handling: enables cookie-based sessions
		AllowCredentials: true,
		//browser caches preflight responses for 5 min; performance: reduces OPTIONS requests
		MaxAge: 300,
	}))

	// Static File Serving
	//web.Files: embedded filesystem from cmd/web/efs.go
	//http.FS(): converts to http.FileSystem interface
	//http.FileServer(): created handlers for serving files
	//echo.WrapHandler(): adapts http.Handler to echo handler
	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))

	//templ.Handler, convert templ component function, web.HelloForm(), to http.Handler
	e.GET("/web", echo.WrapHandler(templ.Handler(web.HelloForm())))           //web page route
	e.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler))) //api route

	//api route
	e.GET("/", s.HelloWorldHandler)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error { //s.HelloWorldHandler is a meathod on Server struct
	resp := map[string]string{
		"message": "Hello World",
	}

	//c.JSON() marshalls resp to JSON and sets the Content-Type
	//pass explicit HTTP status
	return c.JSON(http.StatusOK, resp)
}

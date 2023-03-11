package httpserver

import (
	"net/http"

	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) AddSwagger() {
	sh := http.StripPrefix("/swagger/", http.FileServer(http.Dir(s.config.SwaggerPath)))

	s.server.GET("/swagger/*", echo.WrapHandler(sh))
}

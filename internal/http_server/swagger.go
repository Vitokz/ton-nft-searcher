package httpserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) AddSwagger() {
	sh := http.StripPrefix("/swagger/", http.FileServer(http.Dir(s.config.SwaggerPath)))

	s.server.GET("/swagger/*", echo.WrapHandler(sh))
}

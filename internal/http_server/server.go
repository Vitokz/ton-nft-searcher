package httpserver

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/Vitokz/ton-nft-searcher/internal/protocol"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

type Config struct {
	Host        string
	Port        string
	SwaggerPath string
}

type UseCases interface {
	SearchUserNFTs(ctx context.Context, req protocol.NFTSearchReq) (protocol.NFTSearchResp, error)
}

type Server struct {
	config Config

	server   *echo.Echo
	useCases UseCases

	logger zerolog.Logger
}

func New(config Config, useCases UseCases, logger zerolog.Logger) *Server {
	e := echo.New()
	e.Use(middleware.Logger())

	logger = logger.With().Str("module", "http-server").Logger()

	return &Server{
		config:   config,
		server:   e,
		useCases: useCases,
		logger:   logger,
	}
}

func (s *Server) Start() {
	s.SetRoutes()

	err := s.server.Start(net.JoinHostPort(s.config.Host, s.config.Port))
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			s.logger.Warn().Err(err).Msg("Failed to start http server")
		}
	}
}

func (s *Server) Stop() {
	err := s.server.Shutdown(context.Background())
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to shutdown server")
	}
}

package searcher

import (
	"github.com/Vitokz/ton-nft-searcher/internal/config"
	httpserver "github.com/Vitokz/ton-nft-searcher/internal/http_server"
	"github.com/Vitokz/ton-nft-searcher/internal/logger"
	tonclient "github.com/Vitokz/ton-nft-searcher/internal/ton_client"
	"github.com/Vitokz/ton-nft-searcher/internal/usecases"
	"github.com/spf13/viper"
)

const serviceName = "ton_searcher"

type Searcher struct {
	server   *httpserver.Server
	usecases *usecases.Service
	ton      *tonclient.Client
}

func New() (*Searcher, error) {
	err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	log, err := logger.NewWithDefaultWriters(
		logger.WithModuleName(serviceName),
	)
	if err != nil {
		return nil, err
	}

	client, err := tonclient.New(viper.GetString("SEARCHER_TON_BASE_URL"), log)
	if err != nil {
		return nil, err
	}

	useCases := usecases.New(client, log)

	server := httpserver.New(httpserver.Config{
		Host:        viper.GetString("SEARCHER_HOST"),
		Port:        viper.GetString("SEARCHER_PORT"),
		SwaggerPath: viper.GetString("SEARCHER_SWAGGER_PATH"),
	}, useCases, log)

	return &Searcher{
		server:   server,
		usecases: useCases,
		ton:      client,
	}, nil
}

func (s *Searcher) Start() {
	s.server.Start()
}

func (s *Searcher) Stop() {
	s.server.Stop()
}

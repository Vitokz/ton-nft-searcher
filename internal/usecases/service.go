package usecases

import (
	"github.com/Vitokz/ton-nft-searcher/internal/protocol"
	"github.com/Vitokz/ton-nft-searcher/internal/ton_client/dto"
	"github.com/rs/zerolog"
)

type Ton interface {
	NFTSearch(args dto.NFTSearchArqs) (protocol.NFTSearchResp, error)
}

type Service struct {
	ton    Ton
	logger zerolog.Logger
}

func New(ton Ton, logger zerolog.Logger) *Service {
	logger = logger.With().Str("module", "use-cases").Logger()

	return &Service{
		ton:    ton,
		logger: logger,
	}
}

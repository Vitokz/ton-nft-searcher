package usecases

import (
	"context"

	"github.com/Vitokz/ton-nft-searcher/internal/protocol"
)

func (s *Service) SearchUserNFTs(ctx context.Context, req protocol.NFTSearchReq) (protocol.NFTSearchResp, error) {
	resp, err := s.ton.NFTSearch(req)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to search user NFTs")

		return protocol.NFTSearchResp{}, err
	}

	return resp, nil
}

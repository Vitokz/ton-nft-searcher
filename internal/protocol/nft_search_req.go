package protocol

import (
	"github.com/Vitokz/ton-nft-searcher/internal/entities"
	"github.com/Vitokz/ton-nft-searcher/internal/ton_client/dto"
)

type NFTSearchReq = dto.NFTSearchArqs

type NFTSearchResp struct {
	NFTs []entities.NFTItem
}

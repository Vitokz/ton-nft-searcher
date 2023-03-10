package dto

import (
	"github.com/Vitokz/ton-nft-searcher/internal/ton_client/entities"
	"github.com/tonkeeper/tongo"
)

type NFTSearchArqs struct {
	OwnerAddr      *tongo.AccountID
	CollectionAddr *tongo.AccountID
	IncludeOnSale  *bool
	Limit          int
	Offset         int
}

type NFTSearchResp struct {
	NFTItems entities.NFTItems `json:"nft_items"`
}

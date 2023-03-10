package transport

import "github.com/Vitokz/ton-nft-searcher/internal/http_server/types"

// UserNFTsResponse is response to user nfts request
// swagger:model
type UserNFTsResponse struct {
	NFTItems types.NFTItems `json:"nft_items"`
}

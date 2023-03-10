package types

import (
	"github.com/Vitokz/ton-nft-searcher/internal/entities"
)

// swagger:model
type NFTItem struct {
	Address           string     `json:"address"`
	ApprovedBy        []string   `json:"approved_by"`
	Collection        Collection `json:"collection"`
	CollectionAddress string     `json:"collection_address"`
	Index             int        `json:"index"`
	Metadata          Metadata   `json:"metadata"`
	NFTOwner          NFTOwner   `json:"owner"`
	Previews          Previews   `json:"previews"`
	Verified          bool       `json:"verified"`
}

func (n *NFTItem) FromDTO(dto entities.NFTItem) {
	var collection Collection
	collection.FromDTO(dto.Collection)

	var metadata Metadata
	metadata.FromDTO(dto.Metadata)

	var nftOwner NFTOwner
	nftOwner.FromDTO(dto.NFTOwner)

	var previews Previews
	previews.FromDTO(dto.Previews)

	*n = NFTItem{
		Address:           dto.Address.ToHuman(true, false),
		ApprovedBy:        dto.ApprovedBy,
		Collection:        collection,
		CollectionAddress: collection.Address,
		Index:             dto.Index,
		Metadata:          metadata,
		NFTOwner:          nftOwner,
		Previews:          previews,
		Verified:          dto.Verified,
	}
}

// swagger:model
type NFTItems []NFTItem

func (n *NFTItems) FromDTO(dto entities.NFTItems) {
	nfts := make([]NFTItem, len(dto))

	for i, v := range dto {
		var nft NFTItem
		nft.FromDTO(v)

		nfts[i] = nft
	}

	*n = nfts
}

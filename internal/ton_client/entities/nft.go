package entities

import (
	"github.com/Vitokz/ton-nft-searcher/internal/entities"
	"github.com/tonkeeper/tongo"
)

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

func (n *NFTItem) ToDTO() entities.NFTItem {
	nftAddr := tongo.MustParseAccountID(n.Address)

	collection := n.Collection.ToDTO()

	return entities.NFTItem{
		Address:           nftAddr,
		ApprovedBy:        n.ApprovedBy,
		Collection:        collection,
		CollectionAddress: collection.Address,
		Index:             n.Index,
		Metadata:          n.Metadata.ToDTO(),
		NFTOwner:          n.NFTOwner.ToDTO(),
		Previews:          n.Previews.ToDTO(),
		Verified:          n.Verified,
	}
}

type NFTItems []NFTItem

func (n NFTItems) ToDTO() entities.NFTItems {
	nfts := make([]entities.NFTItem, len(n))

	for i, v := range n {
		nfts[i] = v.ToDTO()
	}

	return nfts
}

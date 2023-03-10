package entities

import "github.com/tonkeeper/tongo"

type NFTItem struct {
	Address           tongo.AccountID
	ApprovedBy        []string
	Collection        Collection
	CollectionAddress tongo.AccountID
	Index             int
	Metadata          Metadata
	NFTOwner          NFTOwner
	Previews          []Preview
	Verified          bool
}

type NFTItems []NFTItem

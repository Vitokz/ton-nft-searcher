package entities

import (
	"github.com/Vitokz/ton-nft-searcher/internal/entities"
	"github.com/tonkeeper/tongo"
)

type Collection struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

func (c *Collection) ToDTO() entities.Collection {
	addr := tongo.MustParseAccountID(c.Address)

	return entities.Collection{
		Address: addr,
		Name:    c.Name,
	}
}

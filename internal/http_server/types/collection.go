package types

import (
	"github.com/Vitokz/ton-nft-searcher/internal/entities"
)

// swagger:model
type Collection struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

func (c *Collection) FromDTO(collection entities.Collection) {
	*c = Collection{
		Address: collection.Address.ToHuman(true, false),
		Name:    collection.Name,
	}
}

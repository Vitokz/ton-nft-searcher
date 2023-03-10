package types

import (
	"github.com/Vitokz/ton-nft-searcher/internal/entities"
)

// Metadata is NFT item meta
// swagger:model
type Metadata struct {
	Image      string     `json:"image"`
	Name       string     `json:"name"`
	Attributes Attributes `json:"attributes"`
}

func (m *Metadata) FromDTO(dto entities.Metadata) {
	var attributes Attributes
	attributes.FromDTO(dto.Attributes)

	*m = Metadata{
		Image:      dto.Image,
		Name:       dto.Name,
		Attributes: attributes,
	}
}

// Attributes is array of NFT attributes
// swagger:model
type Attributes []Attribute

func (a *Attributes) FromDTO(dto entities.Attributes) {
	attrs := make([]Attribute, len(dto))

	for i, v := range dto {
		var attr Attribute
		attr.FromDTO(v)

		attrs[i] = attr
	}

	*a = attrs
}

// Attribute is nft attribute
// swagger:model
type Attribute struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
}

func (a *Attribute) FromDTO(dto entities.Attribute) {
	*a = Attribute{
		TraitType: dto.TraitType,
		Value:     dto.Value,
	}
}

// NFTOwner is owner of nft
// swagger:model
type NFTOwner struct {
	Address string `json:"address"`
	IsScam  bool   `json:"is_scam"`
}

func (n *NFTOwner) FromDTO(dto entities.NFTOwner) {
	*n = NFTOwner{
		Address: dto.Address.ToHuman(true, false),
		IsScam:  dto.IsScam,
	}
}

// Previews is array of previews
// swagger:model
type Previews []Preview

func (p *Previews) FromDTO(dto entities.Previews) {
	previews := make(Previews, len(dto))

	for i, v := range dto {
		var preview Preview
		preview.FromDTO(v)

		previews[i] = preview
	}

	*p = previews
}

// Preview if NFT preview
// swagger:model
type Preview struct {
	Resolution string `json:"resolution"`
	URL        string `json:"url"`
}

func (p *Preview) FromDTO(dto entities.Preview) {
	*p = Preview{
		Resolution: dto.Resolution,
		URL:        dto.URL,
	}
}

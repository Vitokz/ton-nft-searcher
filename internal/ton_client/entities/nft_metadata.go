package entities

import (
	"github.com/Vitokz/ton-nft-searcher/internal/entities"
	"github.com/tonkeeper/tongo"
)

type Metadata struct {
	Image      string `json:"image"`
	Name       string `json:"name"`
	Attributes Attributes
}

func (m *Metadata) ToDTO() entities.Metadata {
	attributes := m.Attributes

	return entities.Metadata{
		Image:      m.Image,
		Name:       m.Name,
		Attributes: attributes.ToDTO(),
	}
}

type Attributes []Attribute

func (a Attributes) ToDTO() entities.Attributes {
	attrs := make([]entities.Attribute, len(a))
	for i, v := range a {
		attrs[i] = v.ToDTO()
	}

	return attrs
}

type Attribute struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
}

func (a *Attribute) ToDTO() entities.Attribute {
	return entities.Attribute{
		TraitType: a.TraitType,
		Value:     a.Value,
	}
}

type NFTOwner struct {
	Address string `json:"address"`
	IsScam  bool   `json:"is_scam"`
}

func (n *NFTOwner) ToDTO() entities.NFTOwner {
	addr := tongo.MustParseAccountID(n.Address)

	return entities.NFTOwner{
		Address: addr,
		IsScam:  n.IsScam,
	}
}

type Previews []Preview

func (p Previews) ToDTO() entities.Previews {
	previews := make([]entities.Preview, len(p))
	for i, v := range p {
		previews[i] = v.ToDTO()
	}

	return previews
}

type Preview struct {
	Resolution string `json:"resolution"`
	URL        string `json:"url"`
}

func (p Preview) ToDTO() entities.Preview {
	return entities.Preview{
		Resolution: p.Resolution,
		URL:        p.URL,
	}
}

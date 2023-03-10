package entities

import "github.com/tonkeeper/tongo"

type Metadata struct {
	Image      string
	Name       string
	Attributes []Attribute
}

type Attributes []Attribute

type Attribute struct {
	TraitType string
	Value     interface{}
}

type NFTOwner struct {
	Address tongo.AccountID
	IsScam  bool
}

type Previews []Preview

type Preview struct {
	Resolution string
	URL        string
}

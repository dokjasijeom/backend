package model

type ProviderLinkModel struct {
	ProviderId uint   `json:"providerId,omitempty"`
	Link       string `json:"link,omitempty"`
}

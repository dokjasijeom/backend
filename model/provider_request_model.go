package model

type ProviderRequestModel struct {
	Id          uint   `json:"id,omitempty"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Description string `json:"description,omitempty"`
	HomepageUrl string `json:"homepageUrl,omitempty"`
}

package model

type PublisherRequestModel struct {
	Id          uint   `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	HomepageUrl string `json:"homepageUrl,omitempty"`
}

package model

type SeriesModel struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	ISBN        string `json:"isbn"`
	ECNNumber   string `json:"ecn_number"`
	SeriesType  string `json:"series_type"`
}

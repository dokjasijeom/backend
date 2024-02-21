package model

type SeriesRequestModel struct {
	Title       string `json:"title,omitempty" validate:"required"`
	Description string `json:"description,omitempty"`
	PublishDay  string `json:"publishDay,omitempty" validate:"oneof=MON TUE WED THU FRI SAT SUN"`
	SeriesType  string `json:"seriesType,omitempty" validate:"oneof=webnovel webtoon"`
}

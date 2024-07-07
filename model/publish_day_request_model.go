package model

type PublishDayRequestModel struct {
	Day          string `json:"day" validate:"required, oneof=MON TUE WED THU FRI SAT SUN"`
	DisplayDay   string `json:"displayDay" validate:"required, oneOf=월 화 수 목 금 토 일"`
	DisplayOrder uint   `json:"displayOrder"`
}

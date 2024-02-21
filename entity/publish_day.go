package entity

type PublishDay struct {
	Id           uint   `gorm:"primaryKey;column:id;type:int(11);autoIncrement;not null" json:"id,omitempty"`
	Day          string `gorm:"column:day;type:varchar(10);unique;not null" json:"day"`
	DisplayDay   string `gorm:"column:display_day;type:varchar(10);unique;not null" json:"displayDay"`
	DisplayOrder uint   `gorm:"column:display_order;type:int(11);null" json:"displayOrder,omitempty"`
	// Day: 월, 화, 수, 목, 금, 토, 일, 자유연재, 주7회, 완결, 휴재, 월~금 // '월,수,금' 화,목, 토, 일,
	//DisplayDay: 월요일, 화요일, 수요일, 목요일, 금요일, 토요일, 일요일, 자유연재, 주 7회, 완결, 휴재, 월~금
	//DisplayOrder: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
}

package entity

type PublishDay struct {
	Id  uint   `gorm:"primaryKey;column:id;type:int(11)"`
	Day string `gorm:"column:day;type:varchar(10);unique;not null"`
	// Day: 월, 화, 수, 목, 금, 토, 일, 자유연재, 주7회, 완결, 휴재, 월~금 // '월,수,금' 화,목, 토, 일,
}

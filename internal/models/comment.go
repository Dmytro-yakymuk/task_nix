package models

type Comment struct {
	PostId int    `gorm:"not null" json:"postId" xml:"postId"`
	UserId string `gorm:"type:varchar(255); not null" json:"userId" xml:"userId"`
	Id     int    `gorm:"autoIncrement:false; not null; unique" json:"id" xml:"id"`
	Name   string `gorm:"type:varchar(255); not null" json:"name" xml:"name"`
	Email  string `gorm:"type:varchar(255); not null" json:"email" xml:"email"`
	Body   string `gorm:"type:text; not null" json:"body" xml:"body"`
}

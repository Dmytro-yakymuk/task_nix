package models

type Comment struct {
	PostId int    `gorm:"not null" json:"postId"`
	Id     int    `gorm:"autoIncrement:false; not null; unique" json:"id"`
	Name   string `gorm:"type:varchar(255); not null" json:"name"`
	Email  string `gorm:"type:varchar(255); not null" json:"email"`
	Body   string `gorm:"type:text; not null" json:"body"`
}

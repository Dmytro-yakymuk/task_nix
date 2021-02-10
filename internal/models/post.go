package models

type Post struct {
	UserId  string    `gorm:"type:varchar(255);not null" json:"userId" xml:"userId"`
	Id      int       `gorm:"autoIncrement:false; not null; unique" json:"id" xml:"id"`
	Title   string    `gorm:"type:varchar(255); not null" json:"title" xml:"title"`
	Body    string    `gorm:"type:text; not null" json:"body" xml:"body"`
	Comment []Comment `gorm:"ForeignKey:PostId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
}

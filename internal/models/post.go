package models

type Post struct {
	UserId  int       `gorm:"not null" json:"userId"`
	Id      int       `gorm:"autoIncrement:false; not null; unique" json:"id"`
	Title   string    `gorm:"type:varchar(255); not null" json:"title"`
	Body    string    `gorm:"type:text; not null" json:"body"`
	Comment []Comment `gorm:"ForeignKey:PostId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments"`
}

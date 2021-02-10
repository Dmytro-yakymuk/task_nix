package models

type User struct {
	Id      string    `gorm:"type:varchar(255); not null; unique" json:"id"`
	Name    string    `gorm:"type:varchar(255); not null" json:"name"`
	Email   string    `gorm:"type:varchar(255); not null; unique" json:"email"`
	Post    []Post    `gorm:"ForeignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
	Comment []Comment `gorm:"ForeignKey:UserId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" xml:"-"`
}

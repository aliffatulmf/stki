package model

type Corpus struct {
	ID       uint   `form:"id,omitempty" json:"id" gorm:"primaryKey;autoIncrement"`
	Title    string `form:"title" json:"title"`
	Body     string `form:"content" json:"content"`
	Document string `form:"document" json:"document" gorm:"uniqueIndex"`
}

package model

type Corpus struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Document string `json:"document"`
}
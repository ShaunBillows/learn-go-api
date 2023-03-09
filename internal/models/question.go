package models

type Question struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	Topic            string `gorm:"not null" json:"topic"`
	Question         string `gorm:"not null;unique" json:"question"`
	Answer           string `gorm:"not null" json:"answer"`
	IncorrectAnswer1 string `gorm:"not null" json:"incorrect_answer_1"`
	IncorrectAnswer2 string `gorm:"not null" json:"incorrect_answer_2"`
	Difficulty       uint   `gorm:"not null" json:"difficulty"`
}

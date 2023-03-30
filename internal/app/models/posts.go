package models

type Posts struct {
	id         string  `gorm:"primaryKey;autoIncrement"`
	text       string  `gorm:"text"`
	contentUrl string  `gorm:"text"`
	authorId   []Users `gorm:"foreignKey:id"` //?
}

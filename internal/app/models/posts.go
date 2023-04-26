package models

type Posts struct {
	ID         string `gorm:"primaryKey;autoIncrement"`
	Text       string `gorm:"text"`
	ContentUrl string `gorm:"text"`
	AuthorId   uint
	User       Users `gorm:"-:migration;foreignKey:AuthorId"`
}

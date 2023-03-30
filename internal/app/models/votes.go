package models

type Votes struct {
	id       uint    `gorm:"primaryKey;autoIncrement"`
	userId   []Users `gorm:"foreignKey:id"`
	postId   []Posts `gorm:"foreignKey:id"`
	typeVote bool
}

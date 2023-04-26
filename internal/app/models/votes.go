package models

type Votes struct {
	ID       uint  `gorm:"primaryKey;autoIncrement"`
	UserId   uint  `gorm:"foreignKey:UserID"`
	PostId   uint  `gorm:"foreignKey:PostID"`
	TypeVote bool  `gorm:"boolean;index"`
	User     Users `gorm:"foreignKey:UserId"`
	Post     Posts `gorm:"foreignKey:PostId"`
}

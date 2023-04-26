package models

type Subscriptions struct {
	ID           uint  `gorm:"primaryKey;autoIncrement"`
	SubscriberId uint  `gorm:"foreignKey:UserID"`
	AuthorId     uint  `gorm:"foreignKey:UserID"`
	Subscriber   Users `gorm:"foreignKey:SubscriberId"`
	Author       Users `gorm:"foreignKey:AuthorId"`
}

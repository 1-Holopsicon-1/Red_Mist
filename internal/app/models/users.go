package models

type Users struct {
	ID         uint    `gorm:"primaryKey;autoIncrement"`
	Username   string  `gorm:"unique;not null"`
	FirstName  string  `gorm:"varchar(120)"`
	SecondName string  `gorm:"varchar(120)"`
	Email      string  `gorm:"varchar;unique"`
	Password   string  `gorm:"text"`
	Role       string  `gorm:"not null"`
	Posts      []Posts `gorm:"constraint:OnDelete:CASCADE; foreignKey:AuthorId;"`
}

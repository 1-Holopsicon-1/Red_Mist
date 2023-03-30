package models

type Users struct {
	id         uint   `gorm:"primaryKey;autoIncrement"`
	firstName  string `gorm:"varchar(120)"`
	secondName string `gorm:"varchar(120)"`
	email      string `gorm:"varchar;unique"`
}

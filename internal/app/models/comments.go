package models

import "github.com/jackc/pgtype"

type Comments struct {
	ID     uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId int         `json:"user_id"`
	PostId int         `json:"post_id"`
	Date   pgtype.Date `json:"date"`
}

package model

type Order struct {
	ID int64 `gorm:"primary_key:auto_increment" json:"id"`
}

package model

import "time"

type Order struct {
	ID            uint64    `gorm:"primary_key:auto_increment" json:"id"`
	IDProduct     uint64    `gorm:"type:int" json:"id_product"`
	JumlahProduct uint64    `gorm:"type:int" json:"jumlah_product"`
	TotalHarga    uint64    `gorm:"type:int" json:"total_harga"`
	Alamat        string    `gorm:"type:varchar(255)" json:"alamat"`
	Status        string    `gorm:"type:varchar(255)" json:"status"`
	IDUser        uint64    `gorm:"type:int" json:"id_user"`
	CreatedAt     time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"type:datetime" json:"updated_at"`
}

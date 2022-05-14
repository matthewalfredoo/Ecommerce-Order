package dto

type NewOrderDTO struct {
	IDProduct     uint64 `json:"id_product" form:"id_product" binding:"required"`
	JumlahProduct uint64 `json:"jumlah_product" form:"jumlah_product" binding:"required"`
	Harga         uint64 `json:"harga" form:"harga"`
	Alamat        string `json:"alamat" form:"alamat" binding:"required"`
	IDUser        uint64 `json:"id_user" form:"id_user" binding:"required"`
}

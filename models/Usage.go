package models

type Usage struct {
    IdPemakaian int      `gorm:"column:id_pemakaian;" json:"id_pemakaian"`
    AssetCode   string   `gorm:"column:kode_aset;type:varchar(20);" json:"kode_aset"`
    Inventory   Inventory `gorm:"foreignKey:AssetCode;references:AssetCode" json:"Inventory"`
    IdRuangan   string   `gorm:"column:id_ruangan;" json:"id_ruangan"`
    Room        Room     `gorm:"foreignKey:IdRuangan;references:IdRuangan" json:"Room"`
    CreatedAt   string   `gorm:"column:created_at" json:"created_at"`
    UpdatedAt   string   `gorm:"column:updated_at" json:"updated_at"`
}



func (Usage) TableName() string {
    return "pemakaian"
}

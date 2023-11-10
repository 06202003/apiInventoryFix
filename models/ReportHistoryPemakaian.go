package models



type ReportHistoryPemakaian struct {
	IdHistory               int64       `gorm:"column:id;" json:"id"`
	OldEmployeeID    string    `gorm:"column:nomor_induk_old;type:varchar(45)" json:"nomor_induk_old"`
	NewEmployeeID    string    `gorm:"column:nomor_induk_new;type:varchar(45)" json:"nomor_induk_new"`
	UsageDate        string`gorm:"column:tanggal" json:"tanggal"`
	OldRoom          string    `gorm:"column:ruangan_old;type:varchar(20)" json:"ruangan_old"`
	NewRoom          string    `gorm:"column:ruangan_new;type:varchar(20)" json:"ruangan_new"`
	CreatedAt        string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        string `gorm:"column:updated_at" json:"updated_at"`
	
	AssetCode 		 string `gorm:"column:kode_aset;type:varchar(20);primaryKey" json:"kode_aset"`
    Inventory  Inventory `gorm:"foreignKey:AssetCode" json:"Inventory"`
}

func (ReportHistoryPemakaian) TableName() string {
	return "history_pemakaian"
}


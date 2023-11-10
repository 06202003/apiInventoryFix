package models

type ReportHistoryPerbaikan struct {
	ID                   int64       `gorm:"column:id;primaryKey" json:"id"`
	RepairDate           string `gorm:"column:tanggal_perbaikan" json:"tanggal_perbaikan"`
	Cost                 int64       `gorm:"column:biaya" json:"biaya"`
	Description          string    `gorm:"column:deskripsi;type:varchar(255)" json:"deskripsi"`
	DamageDate           string `gorm:"column:tanggal_kerusakan" json:"tanggal_kerusakan"`
	RepairCompletionDate string `gorm:"column:tanggal_selesai_perbaikan" json:"tanggal_selesai_perbaikan"`
	CreatedAt            string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt            string `gorm:"column:updated_at" json:"updated_at"`
	AssetCode            string    `gorm:"column:kode_aset;type:varchar(20);primaryKey" json:"kode_aset"`

	Inventory Inventory `gorm:"foreignKey:AssetCode"`
}

func (ReportHistoryPerbaikan) TableName() string {
	return "history_perbaikan"
}

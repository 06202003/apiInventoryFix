package models


type Inventory struct {
	AssetCode      string    `gorm:"column:kode_aset;type:varchar(20);primaryKey" json:"kode_aset"`
	Brand          string    `gorm:"column:merk;type:varchar(45)" json:"merk"`
	Name           string    `gorm:"column:nama;type:varchar(100)" json:"nama"`
	Date           string `gorm:"column:tanggal" json:"tanggal"`
	Price          int64       `gorm:"column:harga" json:"harga"`
	ResidualValue  int64       `gorm:"column:nilai_residu" json:"nilai_residu"`
	UsefulLife     int64       `gorm:"column:masa_manfaat" json:"masa_manfaat"`
	Depreciation   int64       `gorm:"column:depresiasi" json:"depresiasi"`
	Description    string    `gorm:"column:deskripsi;type:varchar(255)" json:"deskripsi"`
	Status         string    `gorm:"column:status;type:varchar(20)" json:"status"`
	RoomID          string    `gorm:"column:id_ruangan;type:varchar(5)" json:"id_ruangan"`
	CreatedAt      string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      string `gorm:"column:updated_at" json:"updated_at"`
	CategoryID     string    `gorm:"column:id_kategori;type:varchar(3);primaryKey" json:"id_kategori"`
	EmployeeID 	   string    `gorm:"column:nomor_induk;type:varchar(20);primaryKey" json:"nomor_induk"`
	Year1		   int64    `gorm:"column:tahun_1" json:"tahun_1"`
	Year2		   int64    `gorm:"column:tahun_2" json:"tahun_2"`
	Year3		   int64    `gorm:"column:tahun_3" json:"tahun_3"`
	Year4		   int64    `gorm:"column:tahun_4" json:"tahun_4"`
	ImageURL	   string   `gorm:"column:img_url;type:varchar(255)" json:"img_url"`
	Vendor		   string   `gorm:"column:vendor;type:varchar(100)" json:"vendor"`

	Category Category `gorm:"foreignKey:CategoryID" json:"Category"`
	Employee Employee `gorm:"foreignKey:EmployeeID" json:"Employee"`
	Room Room  `gorm:"foreignKey:RoomID" json:"Room"`
}

func (Inventory) TableName() string {
	return "inventory"
}

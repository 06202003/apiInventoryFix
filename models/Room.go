package models

type Room struct {
	IdRuangan           string     `gorm:"column:id_ruangan;type:varchar(5)" json:"id_ruangan"`
	Name           		 string 	`gorm:"column:nama" json:"nama"`
	CreatedAt            string 	`gorm:"column:created_at" json:"created_at"`
	UpdatedAt            string 	`gorm:"column:updated_at" json:"updated_at"`

}

func (Room) TableName() string {
	return "ruangan"
}

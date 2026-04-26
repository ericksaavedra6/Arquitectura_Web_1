package repositories

import "gorm.io/gorm"

type ProductoDB struct {
	gorm.Model
	Nombre      string  `gorm:"type:varchar(255);not null"`
	Descripcion string  `gorm:"type:varchar(255);not null"`
	Precio      float64 `gorm:"type:decimal(10,2);not null"`
}

func (ProductoDB) TableName() string {
	return "productos"
}

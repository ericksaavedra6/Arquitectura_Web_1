package repositories

import (
	"entrega1/esaavedra/inventario/internal/core/domain"

	"gorm.io/gorm"
)

type ProductoRepository struct {
	db *gorm.DB
}

func NewProductoRepository(db *gorm.DB) *ProductoRepository {
	return &ProductoRepository{db: db}
}

func (r *ProductoRepository) Save(product *domain.Producto) error {
	dbModel := ProductoDB{
		Nombre:      product.Nombre,
		Descripcion: product.Descripcion,
		Precio:      product.Precio,
	}

	result := r.db.Create(&dbModel)
	if result.Error != nil {
		return result.Error
	}

	product.ID = dbModel.ID
	return nil
}

func (r *ProductoRepository) GetById(id uint) (*domain.Producto, error) {
	var dbModel ProductoDB
	if err := r.db.First(&dbModel, id).Error; err != nil {
		return nil, err
	}

	return &domain.Producto{
		ID:          dbModel.ID,
		Nombre:      dbModel.Nombre,
		Descripcion: dbModel.Descripcion,
		Precio:      dbModel.Precio,
	}, nil
}

func (r *ProductoRepository) Update(product *domain.Producto) error {
	var dbModel ProductoDB
	if err := r.db.First(&dbModel, product.ID).Error; err != nil {
		return err
	}

	dbModel.Nombre = product.Nombre
	dbModel.Descripcion = product.Descripcion
	dbModel.Precio = product.Precio

	return r.db.Save(&dbModel).Error
}

func (r *ProductoRepository) Delete(id uint) error {
	return r.db.Delete(&ProductoDB{}, id).Error
}

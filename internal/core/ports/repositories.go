package ports

import "entrega1/esaavedra/inventario/internal/core/domain"

type ProductoRepository interface {
	Save(product *domain.Producto) error
	GetById(id uint) (*domain.Producto, error)
	Update(product *domain.Producto) error
	Delete(id uint) error
}

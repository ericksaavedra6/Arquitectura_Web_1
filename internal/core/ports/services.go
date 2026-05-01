package ports

import "entrega1/esaavedra/inventario/internal/core/domain"

type ProductService interface {
	Create(product domain.Producto) (domain.Producto, error)
	GetProduct(id uint) (domain.Producto, error)
	Actualizar(producto domain.Producto) (domain.Producto, error)
	Eliminar(id uint) error
}

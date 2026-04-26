package services

import (
	"entrega1/esaavedra/inventario/internal/core/domain"
	"entrega1/esaavedra/inventario/internal/core/ports"
)

type ProductoService struct {
	repo ports.ProductoRepository
}

func NewProductoService(repo ports.ProductoRepository) ports.ProductService {
	return &ProductoService{
		repo: repo,
	}
}

func (service *ProductoService) Create(producto domain.Producto) (domain.Producto, error) {
	err := service.repo.Save(&producto)
	return producto, err
}

func (service *ProductoService) GetProduct(id uint) (domain.Producto, error) {
	prod, err := service.repo.GetById(id)
	if err != nil {
		return domain.Producto{}, err
	}
	return *prod, nil
}

func (service *ProductoService) Actualizar(producto domain.Producto) (domain.Producto, error) {
	err := service.repo.Update(&producto)
	return producto, err
}

func (service *ProductoService) Eliminar(id uint) error {
	err := service.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

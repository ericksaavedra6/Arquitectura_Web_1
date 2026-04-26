package handlers

import (
	"net/http"
	"strconv"

	"entrega1/esaavedra/inventario/internal/core/domain"
	"entrega1/esaavedra/inventario/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type ProductoHandler struct {
	service ports.ProductService
}

func NewProductoHandler(service ports.ProductService) *ProductoHandler {
	return &ProductoHandler{service: service}
}

// POST /productos
func (handler *ProductoHandler) Create(c *gin.Context) {
	var producto domain.Producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := handler.service.Create(producto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, res)
}

// GET /productos/:id
func (handler *ProductoHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := handler.service.GetProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, res)
}

func (handler *ProductoHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// 2. Llamar al servicio (método Eliminar en tu ProductService.go)
	if err := handler.service.Eliminar(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado o ya eliminado"})
		return
	}

	// 3. Respuesta estándar para eliminación exitosa
	c.JSON(http.StatusNoContent, nil)

}

func (handler *ProductoHandler) Update(c *gin.Context) {
	// 1. Obtener ID de la URL
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// 2. Bind del JSON con los nuevos datos
	var p domain.Producto
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	p.ID = uint(id)

	// 3. Llamar al servicio (usa el nombre del método en tu ProductService.go)
	res, err := handler.service.Actualizar(p)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se pudo actualizar el producto"})
		return
	}

	c.JSON(http.StatusOK, res)
}

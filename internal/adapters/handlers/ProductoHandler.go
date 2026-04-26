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

// Create godoc
// @Summary      Crear un nuevo producto
// @Description  Registra un producto en la base de datos usando GORM
// @Tags         productos
// @Accept       json
// @Produce      json
// @Param        producto  body      domain.Producto  true  "Datos del producto"
// @Success      201       {object}  domain.Producto
// @Router       /productos [post]
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

// Create godoc
// @Summary      Consultar un  producto
// @Description  Consultar un producto en la base de datos usando GORM
// @Tags         productos
// @Accept       json
// @Produce      json
// @Param id    path     int true   "id del producto"
// @Success      201       {object}  domain.Producto
// @Router       /productos/{id} [get]
func (handler *ProductoHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	res, err := handler.service.GetProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, res)
}

// Create godoc
// @Summary      Eliminar un  producto
// @Description  Eliminar un producto en la base de datos usando GORM
// @Tags         productos
// @Accept       json
// @Produce      json
// @Param id    path     int true   "id del producto"
// @Success      201       {object}  domain.Producto
// @Router       /productos/{id} [delete]
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

// Create godoc
// @Summary      Eliminar un  producto
// @Description  Eliminar un producto en la base de datos usando GORM
// @Tags         productos
// @Accept       json
// @Produce      json
// @Param id    path     int true   "id del producto"
// @Param        producto  body      domain.Producto  true  "Datos del producto"
// @Success      201       {object}  domain.Producto
// @Router       /productos/{id} [put]
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

package main

import (
	"log"

	"entrega1/esaavedra/inventario/internal/adapters/handlers"
	"entrega1/esaavedra/inventario/internal/adapters/repositories"
	"entrega1/esaavedra/inventario/internal/core/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=db user=admin password=root dbname=inventarios port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a la DB:", err)
	}

	// 2. AutoMigrate (GORM lee tu ProductoDB.go y crea la tabla)
	db.AutoMigrate(&repositories.ProductoDB{})

	// 3. Inyección de Dependencias (Manual / Estilo Go)
	repo := repositories.NewProductoRepository(db)
	service := services.NewProductoService(repo)
	handler := handlers.NewProductoHandler(service)

	// 4. Configuración de Rutas con Gin
	r := gin.Default()
	r.POST("/productos", handler.Create)
	r.DELETE("/productos/:id", handler.Delete)
	r.PUT("/productos/:id", handler.Update)
	r.GET("/productos/:id", handler.Get)

	r.Run()
}

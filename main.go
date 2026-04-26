package main

import (
	"log"

	"entrega1/esaavedra/inventario/internal/adapters/handlers"
	"entrega1/esaavedra/inventario/internal/adapters/repositories"
	"entrega1/esaavedra/inventario/internal/core/services"

	_ "entrega1/esaavedra/inventario/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title           API de Inventarios Erick Saavedra
// @version         1.0
// @description     Servicio para la gestión de productos - Maestría Arquitectura Web.
// @host            localhost:8080
// @BasePath        /api/v1
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
	v1 := r.Group("/api/v1")
	{
		productos := v1.Group("/productos")
		{
			productos.POST("", handler.Create)
			productos.DELETE("/:id", handler.Delete)
			productos.PUT("/:id", handler.Update)
			productos.GET("/:id", handler.Get)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}

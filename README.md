# Arquitectura Web - Entrega 1

## 📝 Descripción de la Entrega
Desarrollo de una **API REST** utilizando **Go** como lenguaje y **Gin** como framework de backend. Se implementó una **Arquitectura Hexagonal** (Ports and Adapters) para garantizar el desacoplamiento de la lógica de negocio, facilitando la mantenibilidad y escalabilidad del sistema.

## 📋 Pre-requisitos
* **Go 1.26.2** instalado en el host.
* **Docker & Docker Compose** para la ejecución orquestada y compilación de imágenes.

---

## 🚀 Puesta en Marcha
Para preparar el entorno y gestionar los módulos de Go, ejecute los siguientes comandos en la raíz del proyecto:

### 1. Inicialización del Módulo
Define el espacio de nombres raíz para las importaciones internas:
```bash
go mod init entrega1/esaavedra/inventario
```
### 2. Gestión de Dependencias
Descarga automáticamente las librerías necesarias (GORM, Gin, Postgres, Swagger) basándose en los imports del código:
```bash
go mod tidy
```
### Verificar que las dependencias estén descargadas localmente
```bash
go mod download
```

## 📖 Documentación con Swagger
La API cuenta con documentación interactiva bajo el estándar OpenAPI. Para regenerar la documentación tras realizar cambios en los Handlers:
```bash
# Instalar la herramienta si no está en el sistema
go install github.com/swaggo/swag/cmd/swag@latest

# Generar los archivos de configuración de OpenAPI (ejecutar en la raíz)
# Nota: Si falla, usa la ruta completa: ~/go/bin/swag init
swag init
```

## 🐳 Ejecución con Docker
Es la forma recomendada para desplegar la arquitectura completa incluyendo la persistencia en PostgreSQL:
La forma recomendada de desplegar la arquitectura completa (API + PostgreSQL):

```bash
# Construir imágenes y levantar servicios en segundo plano
docker-compose up -d --build

# Ver logs de la API para depuración en tiempo real
docker-compose logs -f api

# Detener y eliminar los contenedores y redes
docker-compose down
```


## 🌐 Acceso a la API
Una vez que el contenedor esté arriba, puede interactuar con el servicio a través de:

Swagger UI: http://localhost:8080/swagger/index.html

Base Path de la API: http://localhost:8080/api/v1
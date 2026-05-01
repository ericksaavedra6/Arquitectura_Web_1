# STAGE 1: Compilación
FROM golang:1.26.2-alpine AS builder

# Instalamos git por si alguna dependencia lo requiere
RUN apk add --no-cache git

# Directorio de trabajo
WORKDIR /app

# Copiamos archivos de dependencias primero para aprovechar el cache
COPY go.mod go.sum ./
RUN go mod download

# Copiamos el resto del código
COPY . .

# Compilamos el binario (CGO_ENABLED=0 asegura que sea estático y corra en alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# STAGE 2: Ejecución (Imagen ligera)
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiamos solo el binario desde la etapa anterior
COPY --from=builder /app/main .

# Exponemos el puerto de Gin
EXPOSE 8080

# Comando para arrancar
CMD ["./main"]
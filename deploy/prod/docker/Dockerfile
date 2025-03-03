FROM golang:1.23-alpine AS builder

WORKDIR /app

# Копирование файлов проекта
COPY . .

# Установка зависимостей и сборка
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

# Финальный этап с минимальным Alpine
FROM alpine:latest

WORKDIR /app

# Копирование только бинарного файла из builder
COPY --from=builder /app/main .

# Запуск приложения
CMD ["./main"] 
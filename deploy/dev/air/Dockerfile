FROM golang:1.23-alpine

WORKDIR /app

# Установка необходимых инструментов для разработки
RUN apk add --no-cache git

# Установка air для live-reload
RUN go install github.com/cosmtrek/air@latest

# Копирование файлов проекта
COPY . .

# Установка зависимостей
RUN go mod download

# Команда по умолчанию (будет переопределена в air-compose)
ENTRYPOINT ["air", "-c", "deploy/dev/air/air.toml"]
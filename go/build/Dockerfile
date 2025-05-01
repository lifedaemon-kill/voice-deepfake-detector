# Используем официальный образ Ubuntu 20.04
FROM ubuntu:20.04 AS builder

# Устанавливаем необходимые пакеты
RUN apt-get update && apt-get install -y \
    golang-go \
    && rm -rf /var/lib/apt/lists/*

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o go_telegram_bot ./cmd/main.go

# Используем минимальный образ для запуска
FROM ubuntu:20.04

# Копируем собранное приложение из builder
COPY --from=builder /app/go_telegram_bot /usr/local/bin/go_telegram_bot

# Устанавливаем рабочую директорию
WORKDIR /usr/local/bin

# Запускаем приложение
CMD ["go_telegram_bot"]
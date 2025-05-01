# Используем официальный образ Go
FROM golang:1.24.2

# Устанавливаем рабочую директорию
WORKDIR /go-service

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o run_go_service ./cmd/main.go

# Запускаем приложение
CMD ["./run_go_service"]

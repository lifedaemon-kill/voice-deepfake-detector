FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o go_telegram_bot ./cmd/main.go

CMD ["./go_telegram_bot"]
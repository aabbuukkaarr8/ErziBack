# Stage 1: билд
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Кэшируем модули
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Собираем бинарник из папки cmd/apiserver
RUN go build -o apiserver ./cmd/apiserver

# Stage 2: рантайм
FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /app
# Копируем готовый бинарник
COPY --from=builder /app/apiserver .
COPY --from=builder /app/configs    ./configs

ENV PORT=8080
EXPOSE 8080

CMD ["./apiserver"]

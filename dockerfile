# Dockerfile
FROM golang:1.26-alpine AS builder

# Устанавливаем необходимые инструменты
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Включаем модульный режим
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org

# Скачиваем зависимости (с retry на случай проблем с сетью)
RUN go mod download -x || go mod download

# Копируем все исходники
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
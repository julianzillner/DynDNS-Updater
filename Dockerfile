# syntax=docker/dockerfile:1.3

# Build-Stage
FROM golang:1.22.2 AS builder

WORKDIR /app

# Nur die Mod-Dateien kopieren, um Caching zu nutzen
COPY go.mod go.sum ./
RUN go mod download

# Restlichen Code kopieren
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Run-Stage
FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .

CMD ["./main"]